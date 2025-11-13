package nollywood

import (
	"github.com/NOLLYWOOD-COM/go-sdk/internal/httpclient"
	"github.com/NOLLYWOOD-COM/go-sdk/pkg/catalogue"
	"github.com/NOLLYWOOD-COM/go-sdk/pkg/config"
)

// Client provides access to all Nollywood SDK services
type Client interface {
	// Works returns the work service for catalogue operations
	Works() catalogue.WorkService
}

// client is the concrete implementation of Client
type NollywoodSDKClient struct {
	works      catalogue.WorkService
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

	return &NollywoodSDKClient{
		httpClient: httpClient,
		works:      catalogue.NewWorkService(httpClient),
	}
}

func (c *NollywoodSDKClient) Works() catalogue.WorkService {
	return c.works
}
