package main

import (
	"context"
	"log"
	"net/http"
	"os"
	authcmd "portal_back/authentication/cmd"
	di "portal_back/authentication/cmd"
	companycmd "portal_back/company/cmd"
	documentationDi "portal_back/documentation/impl/di"
	rolesDi "portal_back/roles/impl/di"
)

func InitAppModule() {

	authService, userRequestService, authConn, err := di.InitAuthModule(authcmd.NewConfig())
	defer authConn.Close(context.Background())

	documentConnection := documentationDi.InitDocumentModule(authService)
	defer documentConnection.Close(context.Background())

	// можно инжектить в другие модули
	authService.IsAuthenticated("")

	rolesModule := rolesDi.InitRolesModule()

	companycmd.InitCompanyModule(companycmd.NewConfig(), authService, userRequestService, rolesModule)

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
