package httpclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// GetTokenExpiration extracts the expiration time from a JWT token.
// It parses the token payload and returns the expiration time.
// Returns an error if the token is invalid or doesn't contain an expiration time.
func GetTokenExpiration(token string) (time.Time, error) {
	if token == "" {
		return time.Time{}, fmt.Errorf("token is empty")
	}

	// JWT tokens have three parts: header.payload.signature
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid token format: expected 3 parts, got %d", len(parts))
	}

	// Decode the payload (second part)
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to decode token payload: %w", err)
	}

	// Parse the JSON payload
	var claims struct {
		Exp int64 `json:"exp"` // Expiration time as Unix timestamp
	}

	if err := json.Unmarshal(payload, &claims); err != nil {
		return time.Time{}, fmt.Errorf("failed to parse token claims: %w", err)
	}

	if claims.Exp == 0 {
		return time.Time{}, fmt.Errorf("token does not contain expiration time")
	}

	// Convert Unix timestamp to time.Time
	return time.Unix(claims.Exp, 0), nil
}
