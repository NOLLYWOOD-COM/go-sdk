package catalogue

import "context"

type WorkService interface {
	// GetByIdentifier retrieves a work by its identifier
	GetByIdentifier(ctx context.Context, identifier string) (*Work, error)
	// GetByIdentifiers retrieves multiple works by their identifiers
	GetByIdentifiers(ctx context.Context, identifiers []string) ([]*Work, error)
}

type ArticleService interface {
	// GetByIdentifier retrieves an article by its identifier
	GetByIdentifier(ctx context.Context, identifier string) (*Article, error)
}

type PeopleService interface {
	// GetByIdentifier retrieves a person by its identifier
	GetByIdentifier(ctx context.Context, identifier string) (*Person, error)
	// GetByIdentifiers retrieves multiple people by their identifiers
	GetByIdentifiers(ctx context.Context, identifiers []string) ([]*Person, error)
}
