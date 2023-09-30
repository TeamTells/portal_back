package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"portal_back/api/frontendapi"
	"portal_back/pkg/app/auth"
	"portal_back/pkg/app/token"
	"portal_back/pkg/infrastructure/sql"
	"portal_back/pkg/infrastructure/transport"
)

func main() {

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:12345Q@localhost:5432/teamtells")
	defer conn.Close(context.Background())

	repo := sql.NewTokenStorage(conn)
	tokenService := token.NewService(repo)

	authRepo := sql.NewAuthRepository(conn)
	authService := auth.NewService(authRepo, tokenService)
	server := transport.NewServer(authService)

	http.Handle("/", frontendapi.Handler(server))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
