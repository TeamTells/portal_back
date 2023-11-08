package main

import "portal_back/authentication/impl/di"

func initAppModule() {
	authService := di.InitAuthModule()

	// можно инжектить в другие модули
	authService.IsAuthenticated("")
}
