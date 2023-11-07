package api

import "portal_back/authentication/api/model"

type AuthRequestService interface {
	isAuthenticated(token string) model.AuthValidationResult
}
