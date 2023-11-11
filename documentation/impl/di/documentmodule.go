package di

import (
	"net/http"
	"portal_back/core/network"
	"portal_back/documentation/impl/app/sections"
	"portal_back/documentation/impl/generated/frontendapi"
	"portal_back/documentation/impl/infrastructure/transport"
)

func InitDocumentModule(wrapper network.ResponseWrapper) {
	service := sections.NewSectionService()
	server := transport.NewFrontendServer(service, wrapper)
	http.Handle("/documentation/", frontendapi.Handler(server))
}
