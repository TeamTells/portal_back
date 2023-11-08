package di

import (
	"context"
	"github.com/jackc/pgx/v5"
	"net/http"
	"portal_back/authentication/api"
	"portal_back/authentication/impl/app/auth"
	"portal_back/authentication/impl/app/authrequest"
	"portal_back/authentication/impl/app/token"
	"portal_back/authentication/impl/generated/frontendapi"
	"portal_back/authentication/impl/infrastructure/sql"
	"portal_back/authentication/impl/infrastructure/transport"
)

func InitAuthModule() (api.AuthRequestService, *pgx.Conn) {
	// move const to environment
	conn, _ := pgx.Connect(context.Background(), "postgres://postgres:12345Q@localhost:5432/teamtells")

	repo := sql.NewTokenStorage(conn)
	tokenService := token.NewService(repo)

	authRepo := sql.NewAuthRepository(conn)
	authService := auth.NewService(authRepo, tokenService)
	server := transport.NewServer(authService, tokenService)
	authRequestService := authrequest.NewService()

	http.Handle("/", frontendapi.Handler(server))

	return authRequestService, conn
}
