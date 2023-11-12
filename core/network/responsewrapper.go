package network

import (
	"net/http"
	"portal_back/authentication/api/internalapi"
	"portal_back/authentication/api/internalapi/model"
)

func NewResponseWrapper(authRequestService internalapi.AuthRequestService) ResponseWrapper {
	return ResponseWrapper{authRequestService: authRequestService}
}

type ResponseWrapper struct {
	authRequestService internalapi.AuthRequestService
}

func (wrapper ResponseWrapper) Wrap(w http.ResponseWriter, r *http.Request, block func(RequestInfo)) {
	authResult := wrapper.authRequestService.IsAuthenticated(GetAccessTokenFromHeader(r))
	if authResult == model.NOTFOUND {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if authResult == model.EXPIRED {
		w.WriteHeader(http.StatusUpgradeRequired)
		return
	}

	companyId, err := GetCompanyIdFromHeader(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := GetUserIdFromHeader(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set(ContentTypeHeader, ApplicationJsonType)
	block(RequestInfo{
		CompanyId: companyId,
		UserId:    userId,
	})
}
