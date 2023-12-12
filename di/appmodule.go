package di

import (
	"context"
	"log"
	"net/http"
	"os"
	"portal_back/authentication/impl/di"
	di2 "portal_back/documentation/impl/di"
)

func InitAppModule() {
	authService, authConn := di.InitAuthModule()
	defer authConn.Close(context.Background())

	documentConnection := di2.InitDocumentModule(authService)
	defer documentConnection.Close(context.Background())

	// можно инжектить в другие модули
	authService.IsAuthenticated("")

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	err := http.ListenAndServe(":"+appPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
