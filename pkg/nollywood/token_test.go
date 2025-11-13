package nollywood

import (
	"encoding/base64"
	"encoding/json"
	"testing"
	"time"
)

func TestGetTokenExpiration(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		expectError bool
		checkTime   bool
		expectedExp int64
	}{
		{
			name:        "empty token",
			token:       "",
			expectError: true,
		},
		{
			name:        "invalid format - no dots",
			token:       "invalidtoken",
			expectError: true,
		},
		{
			name:        "invalid format - only one dot",
			token:       "header.payload",
			expectError: true,
		},
		{
			name:        "invalid format - too many parts",
			token:       "header.payload.signature.extra",
			expectError: true,
		},
		{
			name:        "invalid base64 payload",
			token:       "header.!!!invalid!!!.signature",
			expectError: true,
		},
		{
			name:        "invalid JSON payload",
			token:       "header." + base64.RawURLEncoding.EncodeToString([]byte("not json")) + ".signature",
			expectError: true,
		},
		{
			name:        "missing exp claim",
			token:       createTestToken(map[string]interface{}{"sub": "user123"}),
			expectError: true,
		},
		{
			name:        "valid token with exp",
			token:       createTestToken(map[string]interface{}{"exp": int64(1735689600)}),
			expectError: false,
			checkTime:   true,
			expectedExp: 1735689600,
		},
		{
			name:        "valid token with multiple claims",
			token:       createTestToken(map[string]interface{}{"sub": "user123", "exp": int64(1735689600), "iat": int64(1735686000)}),
			expectError: false,
			checkTime:   true,
			expectedExp: 1735689600,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expTime, err := GetTokenExpiration(tt.token)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if tt.checkTime {
					if expTime.Unix() != tt.expectedExp {
						t.Errorf("expected expiration %d, got %d", tt.expectedExp, expTime.Unix())
					}
				}
			}
		})
	}
}

func TestGetTokenExpirationRealJWT(t *testing.T) {
	// Test with a real JWT structure
	// This is a sample JWT with exp claim set to Jan 1, 2025 00:00:00 UTC
	payload := map[string]interface{}{
		"sub":  "1234567890",
		"name": "John Doe",
		"iat":  1516239022,
		"exp":  1735689600, // Jan 1, 2025 00:00:00 UTC
	}

	token := createTestToken(payload)
	expTime, err := GetTokenExpiration(token)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedTime := time.Unix(1735689600, 0)
	if !expTime.Equal(expectedTime) {
		t.Errorf("expected %v, got %v", expectedTime, expTime)
	}
}

// Helper function to create a test JWT token
func createTestToken(claims map[string]interface{}) string {
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}

	headerJSON, _ := json.Marshal(header)
	claimsJSON, _ := json.Marshal(claims)

	headerB64 := base64.RawURLEncoding.EncodeToString(headerJSON)
	claimsB64 := base64.RawURLEncoding.EncodeToString(claimsJSON)

	// We don't need a real signature for testing expiration extraction
	return headerB64 + "." + claimsB64 + ".fakesignature"
}
