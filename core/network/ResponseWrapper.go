package network

import (
	"net/http"
	"portal_back/authentication/api"
	"portal_back/authentication/api/model"
)

func NewResponseWrapper(authRequestService api.AuthRequestService) ResponseWrapper {
	return ResponseWrapper{authRequestService: authRequestService}
}

type ResponseWrapper struct {
	authRequestService api.AuthRequestService
}

func (wrapper ResponseWrapper) Wrap(w http.ResponseWriter, r *http.Request, block func()) {
	authResult := wrapper.authRequestService.IsAuthenticated(GetAccessTokenFromHeader(r))
	if authResult == model.NOTFOUND {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if authResult == model.EXPIRED {
		w.WriteHeader(http.StatusUpgradeRequired)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	block()
}
