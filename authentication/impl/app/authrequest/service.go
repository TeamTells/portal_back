package authrequest

import (
	"portal_back/authentication/api/internalapi"
	"portal_back/authentication/api/internalapi/model"
)

func NewService() internalapi.AuthRequestService {
	return &service{}
}

type service struct {
}

func (s *service) IsAuthenticated(token string) model.AuthValidationResult {
	// TODO: not implemented
	return model.SUCCESS
}
