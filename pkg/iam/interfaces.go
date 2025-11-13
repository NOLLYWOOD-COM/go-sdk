package iam

import "context"

type Authentication interface {
	LoginWithApiKey(ctx context.Context, key string) (TokenPair, error)
	LoginWithPassword(ctx context.Context, identifier, password string) (TokenPair, error)
	RefreshToken(ctx context.Context, refreshToken string) (TokenPair, error)
}
