package catalogue

import (
	"context"
	"fmt"

	"github.com/NOLLYWOOD-COM/go-sdk/internal/httpclient"
)

// NewWorkService creates a new WorkService instance
func NewWorkService(httpClient httpclient.Client) WorkService {
	return &WorkSvc{
		httpClient: httpClient,
	}
}

// GetByIdentifier retrieves a work by its identifier
func (w *WorkSvc) GetByIdentifier(identifier string) (*Work, error) {
	if identifier == "" {
		return nil, fmt.Errorf("identifier cannot be empty")
	}

	url := fmt.Sprintf("%s/works/%s", w.httpClient.GetCatalogueBaseURL(), identifier)
	var work Work

	ctx := context.Background()
	err := w.httpClient.Get(ctx, url, nil, &work)
	if err != nil {
		return nil, fmt.Errorf("failed to get work: %w", err)
	}

	return &work, nil
}
