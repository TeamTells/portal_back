package sql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"portal_back/pkg/app/token"
)

func NewTokenStorage(conn *pgx.Conn) token.Repository {
	return &tokenRepository{conn: conn}
}

type tokenRepository struct {
	conn *pgx.Conn
}

func (r *tokenRepository) GetUserToken(ctx context.Context, userID int) (string, error) {
	query := `
		SELECT token FROM tokens
		WHERE user_id = $1
	`

	var userToken string
	err := r.conn.QueryRow(ctx, query, userID).Scan(&userToken)
	if err == pgx.ErrNoRows {
		return "", token.ErrUserWithTokenNotFound
	} else if err != nil {
		return "", err
	}
	return userToken, nil
}

func (r *tokenRepository) GetUserByToken(ctx context.Context, refreshToken string) (int, error) {
	query := `
		SELECT user_id FROM tokens
		WHERE token = $1
	`

	var userID int
	err := r.conn.QueryRow(ctx, query, refreshToken).Scan(&userID)
	if err == pgx.ErrNoRows {
		return 0, token.ErrUserWithTokenNotFound
	} else if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *tokenRepository) UpdateToken(ctx context.Context, token string, userID int) error {
	query := `
		UPDATE tokens
		SET token = $1
		WHERE user_id = $2
	`
	_, err := r.conn.Exec(ctx, query, token, userID)

	return err
}

func (r *tokenRepository) SaveToken(ctx context.Context, token string, userID int) error {
	query := `
		INSERT INTO tokens (user_id, token)
		VALUES ($1, $2)
	`
	_, err := r.conn.Exec(ctx, query, userID, token)

	return err
}

func (r *tokenRepository) RemoveToken(ctx context.Context, token string) error {
	query := `
		DELETE FROM tokens
		WHERE token = $1
	`
	_, err := r.conn.Exec(ctx, query, token)

	return err
}
