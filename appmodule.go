package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"portal_back/authentication/impl/di"
)

func InitAppModule() {
	authService, authConn := di.InitAuthModule()
	//di2.InitDocumentModule()
	defer authConn.Close(context.Background())

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
