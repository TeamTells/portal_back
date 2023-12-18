package auth

import "context"

type Repository interface {
	GetSaltByLogin(ctx context.Context, login string) (string, error)
	GetPasswordByLogin(ctx context.Context, login string) (string, error)
	GetUserIDByLogin(ctx context.Context, login string) (int, error)
}
