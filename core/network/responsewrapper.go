package network

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"portal_back/authentication/api/internalapi"
	"portal_back/authentication/api/internalapi/model"
)

func Wrap(
	authRequestService internalapi.AuthRequestService,
	w http.ResponseWriter,
	r *http.Request,
	block func(RequestInfo),
) {
	authResult := authRequestService.IsAuthenticated(GetAccessTokenFromHeader(r))
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

func WrapWithBody[TBodyObj any](
	authRequestService internalapi.AuthRequestService,
	w http.ResponseWriter,
	r *http.Request,
	block func(RequestInfo, TBodyObj),
) {
	Wrap(authRequestService, w, r, func(info RequestInfo) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var request TBodyObj
		err = json.Unmarshal(reqBody, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		block(info, request)
	})
}
