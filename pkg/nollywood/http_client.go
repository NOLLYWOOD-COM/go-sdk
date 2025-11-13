package nollywood

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func NewHttpClient(config *Config) HttpClient {
	return &NollywoodHttpClient{
		client: &http.Client{
			Timeout: config.Timeout,
		},
		config: config,
		auth:   &Auth{},
	}
}

func (c *NollywoodHttpClient) GetIAMBaseURL() string {
	return c.config.IAMBaseURL
}

func (c *NollywoodHttpClient) GetCatalogueBaseURL() string {
	return c.config.CatalogueBaseURL
}

func (c *NollywoodHttpClient) Get(ctx context.Context, urlStr string, params map[string]string, result interface{}) error {
	return c.makeRequest(ctx, http.MethodGet, urlStr, params, result, true)
}

func (c *NollywoodHttpClient) Post(ctx context.Context, urlStr string, body interface{}, result interface{}) error {
	return c.makeRequest(ctx, http.MethodPost, urlStr, body, result, true)
}

func (c *NollywoodHttpClient) Put(ctx context.Context, urlStr string, body interface{}, result interface{}) error {
	return c.makeRequest(ctx, http.MethodPut, urlStr, body, result, true)
}

func (c *NollywoodHttpClient) Patch(ctx context.Context, urlStr string, body interface{}, result interface{}) error {
	return c.makeRequest(ctx, http.MethodPatch, urlStr, body, result, true)
}

func (c *NollywoodHttpClient) Delete(ctx context.Context, urlStr string, params map[string]string, result interface{}) error {
	return c.makeRequest(ctx, http.MethodDelete, urlStr, params, result, true)
}

func (c *NollywoodHttpClient) makeRequest(ctx context.Context, method, urlStr string, data interface{}, result interface{}, authenticate bool) error {
	var bodyBytes []byte
	var err error

	// Prepare URL and body based on method
	if method == http.MethodGet || method == http.MethodDelete {
		// Handle query parameters
		if data != nil {
			var queryParams string
			// Check if data is map[string]string or a struct
			if paramsMap, ok := data.(map[string]string); ok {
				// Convert map to query string
				values := url.Values{}
				for k, v := range paramsMap {
					values.Add(k, v)
				}
				queryParams = values.Encode()
			} else {
				// Use StructToQueryParams for structs
				queryParams = StructToQueryParams(data)
			}

			if queryParams != "" {
				// Properly handle existing query parameters
				parsedURL, err := url.Parse(urlStr)
				if err != nil {
					return fmt.Errorf("invalid URL: %w", err)
				}

				if parsedURL.RawQuery != "" {
					parsedURL.RawQuery += "&" + queryParams
				} else {
					parsedURL.RawQuery = queryParams
				}
				urlStr = parsedURL.String()
			}
		}
	} else if data != nil {
		// Prepare JSON body for POST, PUT, PATCH
		bodyBytes, err = json.Marshal(data)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	// Authenticate if required
	if authenticate {
		if err := c.authenticate(ctx); err != nil {
			return fmt.Errorf("authentication failed: %w", err)
		}
	}

	// Execute request with retry logic
	return c.executeWithRetry(ctx, method, urlStr, bodyBytes, result, authenticate)
}

func (c *NollywoodHttpClient) executeWithRetry(ctx context.Context, method, urlStr string, bodyBytes []byte, result interface{}, authenticate bool) error {
	var lastErr error

	for attempt := 0; attempt <= c.config.MaxRetries; attempt++ {
		if attempt > 0 {
			// Wait before retrying with exponential backoff
			time.Sleep(c.config.RetryDelay * time.Duration(attempt))
		}

		// Create fresh request for each attempt
		var body io.Reader
		if len(bodyBytes) > 0 {
			body = bytes.NewReader(bodyBytes)
		}

		req, err := http.NewRequestWithContext(ctx, method, urlStr, body)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		// Set headers
		if len(bodyBytes) > 0 {
			req.Header.Set("Content-Type", "application/json")
		}
		req.Header.Set("User-Agent", c.config.UserAgent)

		// Add authorization header if authenticated
		if authenticate {
			c.authMutex.RLock()
			if c.auth.AccessToken != "" {
				req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
			}
			c.authMutex.RUnlock()
		}

		// Execute request
		resp, err := c.client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("request failed: %w", err)
			continue
		}

		// Handle response
		lastErr = c.handleResponse(resp, result)

		// Close response body immediately
		resp.Body.Close()

		// Check if we should retry
		if lastErr == nil {
			return nil
		}

		// Don't retry on client errors (4xx except 429)
		if resp.StatusCode >= 400 && resp.StatusCode < 500 && resp.StatusCode != http.StatusTooManyRequests {
			return lastErr
		}

		// Retry on server errors (5xx) and rate limiting (429)
		if resp.StatusCode >= 500 || resp.StatusCode == http.StatusTooManyRequests {
			continue
		}

		// For other errors, don't retry
		return lastErr
	}

	return fmt.Errorf("max retries exceeded: %w", lastErr)
}

func (c *NollywoodHttpClient) handleResponse(resp *http.Response, result interface{}) error {
	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check status code
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// Success - unmarshal into result if provided and not 204 No Content
		if result != nil && resp.StatusCode != http.StatusNoContent && len(body) > 0 {
			if err := json.Unmarshal(body, result); err != nil {
				return fmt.Errorf("failed to unmarshal response: %w", err)
			}
		}
		return nil
	}

	// Error response - include body in error message
	if len(body) > 0 {
		return fmt.Errorf("HTTP %d: %s - %s", resp.StatusCode, resp.Status, string(body))
	}
	return fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
}

func (c *NollywoodHttpClient) authenticate(ctx context.Context) error {
	if c.config.ApiKey == "" {
		return fmt.Errorf("no API key provided for authentication")
	}

	c.authMutex.RLock()
	hasValidToken := c.auth.AccessToken != "" && time.Now().Before(c.auth.ExpiresAt)
	hasValidRefresh := c.auth.RefreshToken != "" && time.Now().Before(c.auth.RefreshExpiresAt)
	c.authMutex.RUnlock()

	// If we have a valid access token, we're good
	if hasValidToken {
		return nil
	}

	// Try to refresh token if we have a valid refresh token
	if hasValidRefresh {
		if err := c.refreshToken(ctx); err == nil {
			return nil
		}
		// If refresh fails, fall through to get new token
	}

	// Get new token using API key
	return c.getToken(ctx)
}

func (c *NollywoodHttpClient) getToken(ctx context.Context) error {
	urlStr := fmt.Sprintf("%s/auth/login/key", c.config.IAMBaseURL)
	payload := map[string]string{
		"key": c.config.ApiKey,
	}

	var token TokenPair

	// Make unauthenticated request to avoid infinite recursion
	err := c.makeRequest(ctx, http.MethodPost, urlStr, payload, &token, false)
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	// Parse expiration times from tokens
	accessExpiry, err := GetTokenExpiration(token.AccessToken)
	if err != nil {
		// If we can't parse expiration, set a reasonable default (1 hour)
		accessExpiry = time.Now().Add(1 * time.Hour)
	}

	refreshExpiry, err := GetTokenExpiration(token.RefreshToken)
	if err != nil {
		// If we can't parse expiration, set a reasonable default (7 days)
		refreshExpiry = time.Now().Add(7 * 24 * time.Hour)
	}

	// Update auth state with mutex protection
	c.authMutex.Lock()
	c.auth.AccessToken = token.AccessToken
	c.auth.RefreshToken = token.RefreshToken
	c.auth.ExpiresAt = accessExpiry
	c.auth.RefreshExpiresAt = refreshExpiry
	c.authMutex.Unlock()

	return nil
}

func (c *NollywoodHttpClient) refreshToken(ctx context.Context) error {
	urlStr := fmt.Sprintf("%s/auth/token/refresh", c.config.IAMBaseURL)

	c.authMutex.RLock()
	refreshToken := c.auth.RefreshToken
	c.authMutex.RUnlock()

	payload := map[string]string{
		"refreshToken": refreshToken,
	}

	var token TokenPair

	// Make unauthenticated request to avoid infinite recursion
	err := c.makeRequest(ctx, http.MethodPost, urlStr, payload, &token, false)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	// Parse expiration times from tokens
	accessExpiry, err := GetTokenExpiration(token.AccessToken)
	if err != nil {
		// If we can't parse expiration, set a reasonable default (1 hour)
		accessExpiry = time.Now().Add(1 * time.Hour)
	}

	refreshExpiry, err := GetTokenExpiration(token.RefreshToken)
	if err != nil {
		// If we can't parse expiration, set a reasonable default (7 days)
		refreshExpiry = time.Now().Add(7 * 24 * time.Hour)
	}

	// Update auth state with mutex protection
	c.authMutex.Lock()
	c.auth.AccessToken = token.AccessToken
	c.auth.RefreshToken = token.RefreshToken
	c.auth.ExpiresAt = accessExpiry
	c.auth.RefreshExpiresAt = refreshExpiry
	c.authMutex.Unlock()

	return nil
}
