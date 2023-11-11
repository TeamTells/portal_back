package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"portal_back/authentication/impl/di"
	"portal_back/core/network"
	di2 "portal_back/documentation/impl/di"
)

func InitAppModule() {
	authService, authConn := di.InitAuthModule()
	wrapper := network.NewResponseWrapper(authService)
	di2.InitDocumentModule(wrapper)
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
