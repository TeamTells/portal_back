package token

import "context"

type Repository interface {
	GetUserToken(ctx context.Context, userID int) (string, error)
	UpdateToken(ctx context.Context, token string, userID int) error
	SaveToken(ctx context.Context, token string, userID int) error
}
