package transport

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"portal_back/core/network"
	"portal_back/core/utils"
	"portal_back/documentation/api/frontend"
	"portal_back/documentation/impl/app/sections"
	"portal_back/documentation/impl/domain"
)

func NewFrontendServer(
	sectionService sections.SectionService,
	responseWrapper network.ResponseWrapper,
) frontendapi.ServerInterface {
	return &frontendServer{sectionService: sectionService, responseWrapper: responseWrapper}
}

type frontendServer struct {
	sectionService  sections.SectionService
	responseWrapper network.ResponseWrapper
}

func (server *frontendServer) CreateSection(w http.ResponseWriter, r *http.Request) {
	server.responseWrapper.Wrap(w, r, func(info network.RequestInfo) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var createSectionRequest frontendapi.CreateSectionRequest
		err = json.Unmarshal(reqBody, &createSectionRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		section := domain.Section{
			Id:           domain.NO_ID,
			Title:        *createSectionRequest.Title,
			ThumbnailUrl: *createSectionRequest.ThumbnailUrl,
		}

		server.sectionService.CreateSection(r.Context(), section, info.CompanyId)
	})
}

func (server *frontendServer) GetSections(w http.ResponseWriter, r *http.Request) {
	server.responseWrapper.Wrap(w, r, func(info network.RequestInfo) {
		sections, error := server.sectionService.GetSections(r.Context(), info.CompanyId)
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
		}

		_, error = w.Write(response)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
