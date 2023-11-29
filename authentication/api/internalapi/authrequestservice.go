package internalapi

import (
	"context"
	"portal_back/authentication/api/internalapi/model"
)

type AuthRequestService interface {
	IsAuthenticated(token string) model.AuthValidationResult
	CreateNewUser(ctx context.Context, user model.CreateUserRequest) error
}
