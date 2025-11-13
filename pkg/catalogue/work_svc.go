package catalogue

import (
	"context"
	"fmt"
	"strings"

	"github.com/NOLLYWOOD-COM/go-sdk/internal/httpclient"
)

// NewWorkService creates a new WorkService instance
func NewWorkService(httpClient httpclient.Client) WorkService {
	return &WorkSvc{
		httpClient: httpClient,
	}
}

// GetByIdentifier retrieves a work by its identifier
func (w *WorkSvc) GetByIdentifier(ctx context.Context, identifier string) (*Work, error) {
	if identifier == "" {
		return nil, fmt.Errorf("identifier cannot be empty")
	}

	url := fmt.Sprintf("%s/works/%s", w.httpClient.GetCatalogueBaseURL(), identifier)
	var work Work

	err := w.httpClient.Get(ctx, url, nil, &work)
	if err != nil {
		return nil, fmt.Errorf("failed to get work: %w", err)
	}

	return &work, nil
}

// GetByIdentifiers retrieves multiple works by their identifiers
func (w *WorkSvc) GetByIdentifiers(ctx context.Context, identifiers []string) ([]*Work, error) {
	if len(identifiers) == 0 {
		return nil, fmt.Errorf("identifiers cannot be empty")
	}

	url := fmt.Sprintf("%s/works/batch", w.httpClient.GetCatalogueBaseURL())
	params := map[string]string{
		"identifiers": strings.Join(identifiers, ","),
	}

	var works []*Work

	err := w.httpClient.Get(ctx, url, params, &works)
	if err != nil {
		return nil, fmt.Errorf("failed to get works: %w", err)
	}

	return works, nil
}
