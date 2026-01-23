package httpclient

import "context"

// Client is the internal HTTP client interface for making requests
type Client interface {
	GetIAMBaseURL() string
	GetCatalogueBaseURL() string
	Delete(ctx context.Context, url string, params map[string]string, result interface{}) error
	Get(ctx context.Context, url string, params map[string]string, result interface{}) error
	Patch(ctx context.Context, url string, body interface{}, result interface{}) error
	Post(ctx context.Context, url string, body interface{}, result interface{}) error
	Put(ctx context.Context, url string, body interface{}, result interface{}) error
}

// AuthProvider provides access to authentication tokens.
// This interface allows external components (like GraphQL clients)
// to access the auth system without direct coupling.
type AuthProvider interface {
	// GetAccessToken returns a valid access token.
	// It handles token acquisition and refresh automatically.
	GetAccessToken(ctx context.Context) (string, error)
	// GetUserAgent returns the configured user agent string.
	GetUserAgent() string
}
