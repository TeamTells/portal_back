package di

import (
	"context"
	"log"
	"net/http"
	"os"
	authcmd "portal_back/authentication/cmd"
	di2 "portal_back/documentation/impl/di"
)

func InitAppModule() {
	authService, authConn, err := authcmd.InitAuthModule(authcmd.NewConfig())
	if err != nil {
		log.Fatal("failed init auth module:", err)
	}
	defer authConn.Close(context.Background())

	documentConnection := di2.InitDocumentModule(authService)
	defer documentConnection.Close(context.Background())

	// можно инжектить в другие модули
	authService.IsAuthenticated("")

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	err = http.ListenAndServe(":"+appPort, nil)
	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}

func migrate() {
	err := authcmd.Migrate(authcmd.NewConfig())
	if err != nil {
		log.Fatal("failed migrate auth module:", err)
	}
}
