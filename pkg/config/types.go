package config

import "time"

// Option is a function that modifies a Config
type Option func(*Config)

// Config holds configuration for the Nollywood SDK
type Config struct {
	IAMBaseURL       string        // Base URL for the IAM service
	CatalogueBaseURL string        // Base URL for the Catalogue service
	ApiKey           string        // API key for authentication
	Timeout          time.Duration // Request timeout duration
	RetryDelay       time.Duration // Delay between retries
	MaxRetries       int           // Maximum number of retries for requests
	UserAgent        string        // User-Agent header value
}
