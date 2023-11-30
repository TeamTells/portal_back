package authrequest

import (
	"context"
	"portal_back/authentication/api/internalapi"
	"portal_back/authentication/api/internalapi/model"
)

func NewService() internalapi.AuthRequestService {
	return &service{}
}

type service struct {
}

func (s *service) GetUserIdByEmail(ctx context.Context, email string) (*int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) IsAuthenticated(token string) model.AuthValidationResult {
	// TODO: not implemented
	return model.SUCCESS
}

func (s *service) CreateNewUser(ctx context.Context, user model.CreateUserRequest) error {
	// TODO: not implemented
	return nil
}
