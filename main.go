package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"portal_back/authentication/impl/api/frontendapi"
	"portal_back/authentication/impl/pkg/app/auth"
	"portal_back/authentication/impl/pkg/app/token"
	sql2 "portal_back/authentication/impl/pkg/infrastructure/sql"
	"portal_back/authentication/impl/pkg/infrastructure/transport"
)

func main() {

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:12345Q@localhost:5432/teamtells")
	defer conn.Close(context.Background())

	repo := sql2.NewTokenStorage(conn)
	tokenService := token.NewService(repo)

	authRepo := sql2.NewAuthRepository(conn)
	authService := auth.NewService(authRepo, tokenService)
	server := transport.NewServer(authService, tokenService)

	http.Handle("/", frontendapi.Handler(server))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
