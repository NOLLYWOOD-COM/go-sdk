package catalogue

type WorkService interface {
	GetByIdentifier(identifier string) (*Work, error)
}

type ArticleService interface{}
