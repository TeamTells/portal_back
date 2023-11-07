package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
	"portal_back/authentication/impl/api/frontendapi"
	"portal_back/authentication/impl/pkg/app/auth"
	"portal_back/authentication/impl/pkg/app/token"
	sql2 "portal_back/authentication/impl/pkg/infrastructure/sql"
	"portal_back/authentication/impl/pkg/infrastructure/transport"
)

func main() {

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "app"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", dbUser, dbPassword, dbHost, dbName)

	conn, err := pgx.Connect(context.Background(), connStr)
	defer conn.Close(context.Background())

	repo := sql2.NewTokenStorage(conn)
	tokenService := token.NewService(repo)

	authRepo := sql2.NewAuthRepository(conn)
	authService := auth.NewService(authRepo, tokenService)
	server := transport.NewServer(authService, tokenService)

	http.Handle("/", frontendapi.Handler(server))

	err = http.ListenAndServe(":"+appPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
