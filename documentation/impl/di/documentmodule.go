package di

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
	"portal_back/authentication/api/internalapi"
	frontendapi "portal_back/documentation/api/frontend"
	"portal_back/documentation/impl/app/sections"
	"portal_back/documentation/impl/infrastructure/sql"
	"portal_back/documentation/impl/infrastructure/transport"
)

func InitDocumentModule(authRequestService internalapi.AuthRequestService) *pgx.Conn {
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

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5431/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)
	conn, _ := pgx.Connect(context.Background(), connStr)

	sectionRepository := sql.NewSectionRepository(conn)
	service := sections.NewSectionService(sectionRepository)
	server := transport.NewFrontendServer(service, authRequestService)
	http.Handle("/documentation/", frontendapi.Handler(server))
	return conn
}
