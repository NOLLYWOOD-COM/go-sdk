package catalogue

import (
	"context"
	"fmt"

	"github.com/NOLLYWOOD-COM/go-sdk/internal/httpclient"
)

func NewArticleService(httpClient httpclient.Client) ArticleService {
	return &ArticleSvc{
		httpClient: httpClient,
	}
}

// GetByIdentifier retrieves an article by its identifier
func (a *ArticleSvc) GetByIdentifier(ctx context.Context, identifier string) (*Article, error) {
	if identifier == "" {
		return nil, fmt.Errorf("identifier cannot be empty")
	}

	url := fmt.Sprintf("%s/articles/%s", a.httpClient.GetCatalogueBaseURL(), identifier)
	var article Article

	err := a.httpClient.Get(ctx, url, nil, &article)
	if err != nil {
		return nil, fmt.Errorf("failed to get article: %w", err)
	}

	return &article, nil
}
