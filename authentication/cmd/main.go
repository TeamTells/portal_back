package cmd

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
	"portal_back/authentication/api/frontend"
	"portal_back/authentication/api/internalapi"
	"portal_back/authentication/impl/app/auth"
	"portal_back/authentication/impl/app/authrequest"
	"portal_back/authentication/impl/app/token"
	"portal_back/authentication/impl/app/userrequest"
	"portal_back/authentication/impl/infrastructure/sql"
	"portal_back/authentication/impl/infrastructure/transport"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func InitAuthModule(config Config) (internalapi.AuthRequestService, internalapi.UserRequestService, *pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBName)

	conn, _ := pgx.Connect(context.Background(), connStr)

	repoId := sql.NewTokenStorage(conn)
	tokenService := token.NewService(repoId)

	authRepo := sql.NewAuthRepository(conn)
	authService := auth.NewService(authRepo, tokenService)
	server := transport.NewServer(authService, tokenService)
	authRequestService := authrequest.NewService()
	userRequestService := userrequest.NewService(authService)

	http.Handle("/authorization/", frontendapi.Handler(server))

	return authRequestService, userRequestService, conn, nil
}
