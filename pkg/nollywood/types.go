package nollywood

import (
	"net/http"
	"sync"
	"time"
)

type Option func(*Config)

type Config struct {
	IAMBaseURL       string        // Base URL for the IAM service
	CatalogueBaseURL string        // Base URL for the Catalogue service
	ApiKey           string        // API key for authentication
	Timeout          time.Duration // Request timeout duration
	RetryDelay       time.Duration // Delay between retries
	MaxRetries       int           // Maximum number of retries for requests
	UserAgent        string        // User-Agent header value
}

type NollywoodHttpClient struct {
	auth      *Auth        // Client authentication state
	authMutex sync.RWMutex // Mutex to protect auth state from concurrent access
	client    *http.Client // Underlying HTTP client
	config    *Config      // Configuration for the client
}

type Auth struct {
	AccessToken      string    // Current access token
	RefreshToken     string    // Current refresh token
	ExpiresAt        time.Time // Expiration time of the access token
	RefreshExpiresAt time.Time // Expiration time of the refresh token
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`  // Access token string from server
	RefreshToken string `json:"refreshToken"` // Refresh token string from server
}
