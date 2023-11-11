package di

import (
	"net/http"
	"portal_back/documentation/impl/app/sections"
	"portal_back/documentation/impl/generated/frontendapi"
	"portal_back/documentation/impl/infrastructure/transport"
)

func InitDocumentModule() {
	service := sections.NewSectionService()
	server := transport.NewFrontendServer(service)
	http.Handle("/", frontendapi.Handler(server))
}
