package nollywood

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/NOLLYWOOD-COM/go-sdk/internal/httpclient"
	"github.com/NOLLYWOOD-COM/go-sdk/pkg/catalogue"
	"github.com/NOLLYWOOD-COM/go-sdk/pkg/config"
	nollywoodgql "github.com/NOLLYWOOD-COM/go-sdk/pkg/graphql"
)

// Client provides access to all Nollywood SDK services
type Client interface {
	// Works returns the work service for catalogue operations
	Works() catalogue.WorkService
	// People returns the people service for catalogue operations
	People() catalogue.PeopleService
	// Articles returns the article service for catalogue operations
	Articles() catalogue.ArticleService
	// GraphQL returns the authenticated GraphQL client for custom queries
	GraphQL() graphql.Client
}

// client is the concrete implementation of Client
type NollywoodSDKClient struct {
	works      catalogue.WorkService
	people     catalogue.PeopleService
	articles   catalogue.ArticleService
	graphql    graphql.Client
	httpClient httpclient.Client
}

// NewClient creates a new SDK client with the given configuration
func NewClient(config *config.Config) Client {
	// Convert public Config to internal httpclient.Config
	httpClientConfig := &httpclient.Config{
		IAMBaseURL:       config.IAMBaseURL,
		CatalogueBaseURL: config.CatalogueBaseURL,
		ApiKey:           config.ApiKey,
		Timeout:          config.Timeout,
		RetryDelay:       config.RetryDelay,
		MaxRetries:       config.MaxRetries,
		UserAgent:        config.UserAgent,
	}

	httpClient := httpclient.New(httpClientConfig)

	// Cast httpClient to AuthProvider for GraphQL client
	authProvider := httpClient.(httpclient.AuthProvider)

	return &NollywoodSDKClient{
		httpClient: httpClient,
		works:      catalogue.NewWorkService(httpClient),
		people:     catalogue.NewPeopleService(httpClient),
		articles:   catalogue.NewArticleService(httpClient),
		graphql:    nollywoodgql.NewClient(config.GraphQLEndpoint, authProvider, config.Timeout),
	}
}

func (c *NollywoodSDKClient) Works() catalogue.WorkService {
	return c.works
}

func (c *NollywoodSDKClient) People() catalogue.PeopleService {
	return c.people
}

func (c *NollywoodSDKClient) Articles() catalogue.ArticleService {
	return c.articles
}

func (c *NollywoodSDKClient) GraphQL() graphql.Client {
	return c.graphql
}
