package transport

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"portal_back/api/frontendapi"
	"portal_back/pkg/app/auth"
)

func NewServer(authService auth.Service) frontendapi.ServerInterface {
	return &frontendServer{authService: authService}
}

type frontendServer struct {
	authService auth.Service
}

func (s *frontendServer) GetSaltByLogin(w http.ResponseWriter, r *http.Request, login string) {
	salt, err := s.authService.GetSaltByLogin(r.Context(), login)
	if err == auth.ErrUserNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := json.Marshal(frontendapi.SaltResponse{
		Salt: &salt,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *frontendServer) Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var loginReq frontendapi.LoginRequest
	err = json.Unmarshal(reqBody, &loginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	tokens, err := s.authService.Login(
		r.Context(),
		auth.LoginData{
			Login:    *loginReq.Login,
			Password: *loginReq.Password,
		})

	if err == auth.ErrUserNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else if err == nil {

		resp, err := json.Marshal(frontendapi.TokenResponse{
			AccessJwtToken: &tokens.AccessToken,
			RefreshToken:   &tokens.RefreshToken,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}
