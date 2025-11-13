package httpclient

import (
	"net/http"
	"sync"
	"time"
)

// client is the internal HTTP client implementation
type client struct {
	auth       *AuthState
	authMutex  sync.RWMutex
	httpClient *http.Client
	config     *Config
}

// auth holds authentication state
type AuthState struct {
	AccessToken      string
	RefreshToken     string
	ExpiresAt        time.Time
	RefreshExpiresAt time.Time
}

// TokenPair represents an access/refresh token pair from the server
type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// Config holds configuration for the HTTP client
type Config struct {
	IAMBaseURL       string
	CatalogueBaseURL string
	ApiKey           string
	Timeout          time.Duration
	RetryDelay       time.Duration
	MaxRetries       int
	UserAgent        string
}
