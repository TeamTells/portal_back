package authrequest

import (
	"portal_back/authentication/api"
	"portal_back/authentication/api/model"
)

func NewService() api.AuthRequestService {
	return &service{}
}

type service struct {
}

func (s *service) IsAuthenticated(token string) model.AuthValidationResult {
	// TODO: not implemented
	return model.SUCCESS
}
