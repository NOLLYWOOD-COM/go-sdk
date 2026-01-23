package graphql

import (
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/NOLLYWOOD-COM/go-sdk/internal/httpclient"
)

// NewClient creates a new authenticated GraphQL client.
// The returned client uses the provided AuthProvider for automatic token management.
func NewClient(endpoint string, authProvider httpclient.AuthProvider, timeout time.Duration) graphql.Client {
	// Create authenticated transport that wraps http.DefaultTransport
	transport := httpclient.NewAuthenticatedTransport(authProvider)

	// Create HTTP client with the authenticated transport
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	// Create and return genqlient client
	return graphql.NewClient(endpoint, httpClient)
}
