package di

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
	"portal_back/core/network"
	"portal_back/documentation/impl/app/sections"
	"portal_back/documentation/impl/generated/frontendapi"
	"portal_back/documentation/impl/infrastructure/sql"
	"portal_back/documentation/impl/infrastructure/transport"
)

func InitDocumentModule(wrapper network.ResponseWrapper) *pgx.Conn {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}

	dbName := os.Getenv("DB_DOCUMENTATION_NAME")
	if dbName == "" {
		dbName = "app"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", dbUser, dbPassword, dbHost, dbName)
	conn, _ := pgx.Connect(context.Background(), connStr)

	sectionRepository := sql.NewSectionRepository(conn)
	service := sections.NewSectionService(sectionRepository)
	server := transport.NewFrontendServer(service, wrapper)
	http.Handle("/documentation/", frontendapi.Handler(server))
	return conn
}
