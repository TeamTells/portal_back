package sql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"portal_back/authentication/impl/app/auth"
)

func NewAuthRepository(conn *pgx.Conn) auth.Repository {
	return &repository{conn: conn}
}

type repository struct {
	conn *pgx.Conn
}

func (r *repository) GetSaltByLogin(ctx context.Context, login string) (string, error) {
	query := `
		SELECT salt FROM auth_user 
        WHERE login=$1
	`
	var salt string
	err := r.conn.QueryRow(ctx, query, login).Scan(&salt)
	if err == pgx.ErrNoRows {
		return "", auth.ErrUserNotFound
	} else if err != nil {
		return "", err
	}
	return salt, nil
}

func (r *repository) GetPasswordByLogin(ctx context.Context, login string) (string, error) {
	query := `
		SELECT password FROM auth_user 
        WHERE login=$1
	`
	var password string
	err := r.conn.QueryRow(ctx, query, login).Scan(&password)
	if err == pgx.ErrNoRows {
		return "", auth.ErrUserNotFound
	} else if err != nil {
		return "", err
	}
	return password, nil
}

func (r *repository) GetUserIDByLogin(ctx context.Context, login string) (int, error) {
	query := `
		SELECT id FROM auth_user 
        WHERE login=$1
	`
	var userID int
	err := r.conn.QueryRow(ctx, query, login).Scan(&userID)
	if err == pgx.ErrNoRows {
		return 0, auth.ErrUserNotFound
	} else if err != nil {
		return 0, err
	}
	return userID, nil
}
