package internalapi

import (
	"errors"
)

var UserAlreadyExists = errors.New("user with this email already exists")

type UserRequestService interface {
	CreateNewUser(email string) error
	GetUserId(email string) (int, error)
}
