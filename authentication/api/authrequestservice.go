package api

import "portal_back/authentication/api/model"

type AuthRequestService interface {
	IsAuthenticated(token string) model.AuthValidationResult
}
