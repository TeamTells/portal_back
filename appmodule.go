package main

import (
	"context"
	"log"
	"net/http"
	"portal_back/authentication/impl/di"
)

func InitAppModule() {
	authService, authConn := di.InitAuthModule()
	defer authConn.Close(context.Background())

	// можно инжектить в другие модули
	authService.IsAuthenticated("")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
