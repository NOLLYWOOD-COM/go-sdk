package httpclient

import (
	"net/http"
)

// AuthenticatedTransport is an http.RoundTripper that adds authentication headers.
// It delegates to the existing auth system for token management.
type AuthenticatedTransport struct {
	// Base is the underlying transport to use for requests.
	// If nil, http.DefaultTransport is used.
	Base http.RoundTripper

	// AuthProvider provides access tokens for authentication.
	AuthProvider AuthProvider
}

// RoundTrip implements http.RoundTripper.
// It adds the Authorization header with a Bearer token before delegating to Base.
func (t *AuthenticatedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request to avoid mutating the original
	reqClone := req.Clone(req.Context())

	// Get access token from the auth provider
	token, err := t.AuthProvider.GetAccessToken(req.Context())
	if err != nil {
		return nil, err
	}

	// Set authorization header
	reqClone.Header.Set("Authorization", "Bearer "+token)

	// Set user agent if available
	userAgent := t.AuthProvider.GetUserAgent()
	if userAgent != "" {
		reqClone.Header.Set("User-Agent", userAgent)
	}

	// Use Base transport or default
	transport := t.Base
	if transport == nil {
		transport = http.DefaultTransport
	}

	return transport.RoundTrip(reqClone)
}

// NewAuthenticatedTransport creates a new authenticated transport.
func NewAuthenticatedTransport(authProvider AuthProvider) *AuthenticatedTransport {
	return &AuthenticatedTransport{
		Base:         http.DefaultTransport,
		AuthProvider: authProvider,
	}
}
