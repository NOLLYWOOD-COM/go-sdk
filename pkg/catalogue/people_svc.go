package catalogue

import (
	"context"
	"fmt"
	"strings"

	"github.com/NOLLYWOOD-COM/go-sdk/internal/httpclient"
)

func NewPeopleService(httpClient httpclient.Client) PeopleService {
	return &PeopleSvc{
		httpClient: httpClient,
	}
}

// GetByIdentifier retrieves a person by their identifier
func (p *PeopleSvc) GetByIdentifier(ctx context.Context, identifier string) (*Person, error) {
	if identifier == "" {
		return nil, fmt.Errorf("identifier cannot be empty")
	}

	url := fmt.Sprintf("%s/people/%s", p.httpClient.GetCatalogueBaseURL(), identifier)
	var person Person

	err := p.httpClient.Get(ctx, url, nil, &person)
	if err != nil {
		return nil, fmt.Errorf("failed to get person: %w", err)
	}

	return &person, nil
}

// GetByIdentifiers retrieves multiple people by their identifiers
func (p *PeopleSvc) GetByIdentifiers(ctx context.Context, identifiers []string) ([]*Person, error) {
	if len(identifiers) == 0 {
		return nil, fmt.Errorf("identifiers cannot be empty")
	}

	url := fmt.Sprintf("%s/people/batch", p.httpClient.GetCatalogueBaseURL())
	params := map[string]string{
		"identifiers": strings.Join(identifiers, ","),
	}

	var people []*Person

	err := p.httpClient.Get(ctx, url, params, &people)
	if err != nil {
		return nil, fmt.Errorf("failed to get people: %w", err)
	}

	return people, nil
}
