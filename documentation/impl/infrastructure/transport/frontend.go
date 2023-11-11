package transport

import (
	"encoding/json"
	"net/http"
	"portal_back/core"
	"portal_back/documentation/impl/app/sections"
	"portal_back/documentation/impl/domain"
	"portal_back/documentation/impl/generated/frontendapi"
)

func NewFrontendServer(sectionService sections.SectionService) frontendapi.ServerInterface {
	return &frontendServer{sectionService: sectionService}
}

type frontendServer struct {
	sectionService sections.SectionService
}

func (server *frontendServer) GetSections(w http.ResponseWriter, r *http.Request) {
	sections, error := server.sectionService.GetSections()
	if error != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	sectionJson := core.Map(sections, func(section domain.Section) frontendapi.Section {
		return frontendapi.Section{
			Id:           section.Id,
			ThumbnailUrl: section.ThumbnailUrl,
			Title:        section.Title,
		}
	})

	response, error := json.Marshal(frontendapi.SectionsResponse{Sections: sectionJson})
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, error = w.Write(response)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//
//func (s *frontendServer) GetSaltByLogin(w http.ResponseWriter, r *http.Request, login string) {
//	salt, err := s.authService.GetSaltByLogin(r.Context(), login)
//	if err == auth.ErrUserNotFound {
//		w.WriteHeader(http.StatusNotFound)
//	} else if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//	}
//
//	resp, err := json.Marshal(frontendapi.SaltResponse{
//		Salt: &salt,
//	})
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	_, err = w.Write(resp)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//	}
//}
