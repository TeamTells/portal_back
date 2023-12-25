package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
	authcmd "portal_back/authentication/cmd"
	"portal_back/authentication/impl/di"
	companyDi "portal_back/company/impl/di"
	documentationDi "portal_back/documentation/impl/di"
	rolesDi "portal_back/roles/impl/di"
)

func InitAppModule() {
	conn := createConnection()
	defer conn.Close(context.Background())

	authService, userRequestService := di.InitAuthModule(conn)

	documentConnection := documentationDi.InitDocumentModule(authService)
	defer documentConnection.Close(context.Background())

	// можно инжектить в другие модули
	authService.IsAuthenticated("")

	rolesModule := rolesDi.InitRolesModule()

	companyDi.InitCompanyModule(authService, userRequestService, rolesModule, conn)

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	err := http.ListenAndServe(":"+appPort, nil)
	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}

func createConnection() *pgx.Conn {
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

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", dbUser, dbPassword, dbHost, dbName)

	conn, _ := pgx.Connect(context.Background(), connStr)

	return conn
}

func migrate() {
	err := authcmd.Migrate(authcmd.NewConfig())
	if err != nil {
		log.Fatal("failed migrate auth module:", err)
	}
}
