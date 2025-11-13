package nollywood

import "context"

type HttpClient interface {
	GetIAMBaseURL() string
	GetCatalogueBaseURL() string
	Delete(ctx context.Context, url string, params map[string]string, result interface{}) error
	Get(ctx context.Context, url string, params map[string]string, result interface{}) error
	Patch(ctx context.Context, url string, body interface{}, result interface{}) error
	Post(ctx context.Context, url string, body interface{}, result interface{}) error
	Put(ctx context.Context, url string, body interface{}, result interface{}) error
}
