package transport

import (
	"encoding/json"
	"net/http"
	"portal_back/authentication/api/internalapi"
	"portal_back/core/network"
	"portal_back/core/utils"
	"portal_back/documentation/api/frontend"
	"portal_back/documentation/impl/app/sections"
	"portal_back/documentation/impl/domain"
)

func NewFrontendServer(
	sectionService sections.SectionService,
	authRequestService internalapi.AuthRequestService,
) frontendapi.ServerInterface {
	return &frontendServer{sectionService: sectionService, authRequestService: authRequestService}
}

type frontendServer struct {
	sectionService     sections.SectionService
	authRequestService internalapi.AuthRequestService
}

func (server *frontendServer) GetSections(w http.ResponseWriter, r *http.Request) {
	network.Wrap(server.authRequestService, w, r, func(info network.RequestInfo) {
		sections, error := server.sectionService.GetSections(r.Context(), info.CompanyId, info.UserId)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		sectionJson := utils.Map(sections, func(section domain.Section) frontendapi.Section {
			return frontendapi.Section{
				Id:           section.Id,
				ThumbnailUrl: section.ThumbnailUrl,
				Title:        section.Title,
				IsFavorite:   section.IsFavorite,
			}
		})

		response, error := json.Marshal(frontendapi.SectionsResponse{Sections: sectionJson})
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, error = w.Write(response)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func (server *frontendServer) CreateSection(w http.ResponseWriter, r *http.Request) {
	network.WrapWithBody(server.authRequestService, w, r, func(info network.RequestInfo, request frontendapi.CreateSectionRequest) {
		section := domain.Section{
			Id:           domain.NO_ID,
			Title:        *request.Title,
			ThumbnailUrl: *request.ThumbnailUrl,
		}

		server.sectionService.CreateSection(r.Context(), section, info.CompanyId)
	})
}

func (server *frontendServer) UpdateIsFavoriteSection(w http.ResponseWriter, r *http.Request) {
	network.WrapWithBody(server.authRequestService, w, r, func(info network.RequestInfo, request frontendapi.FavoriteRequest) {
		server.sectionService.UpdateIsFavoriteSection(r.Context(), *request.SectionId, *request.IsFavorite)
	})
}
