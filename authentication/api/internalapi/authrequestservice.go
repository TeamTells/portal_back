package internalapi

import (
	"portal_back/authentication/api/internalapi/model"
)

type AuthRequestService interface {
	IsAuthenticated(token string) model.AuthValidationResult
}
