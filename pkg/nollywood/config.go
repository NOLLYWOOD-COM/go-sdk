package nollywood

import (
	"fmt"
	"time"
)

func WithApiKey(apiKey string) Option {
	return func(c *Config) {
		c.ApiKey = apiKey
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

func WithRetryDelay(retryDelay time.Duration) Option {
	return func(c *Config) {
		c.RetryDelay = retryDelay
	}
}

func WithMaxRetries(maxRetries int) Option {
	return func(c *Config) {
		c.MaxRetries = maxRetries
	}
}

func WithUserAgent(userAgent string) Option {
	return func(c *Config) {
		c.UserAgent = userAgent
	}
}

func WithIAMBaseURL(url string) Option {
	return func(c *Config) {
		c.IAMBaseURL = url
	}
}

func WithCatalogueBaseURL(url string) Option {
	return func(c *Config) {
		c.CatalogueBaseURL = url
	}
}

// DefaultConfig returns a Config struct populated with default values.
func DefaultConfig(iamBaseUrl, catalogueBaseUrl string) *Config {
	return &Config{
		IAMBaseURL:       iamBaseUrl,
		CatalogueBaseURL: catalogueBaseUrl,
		Timeout:          15 * time.Second,
		MaxRetries:       3,
		RetryDelay:       2 * time.Second,
		UserAgent:        fmt.Sprintf("nollywood-go-sdk/%s", SDK_VERSION),
	}
}
