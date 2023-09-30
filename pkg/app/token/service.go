package token

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

var ErrTokenForUserNotFound = errors.New("token for user not found")

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Service interface {
	GenerateTokensForUser(ctx context.Context, userID int) (Tokens, error)
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

type service struct {
	repository Repository
}

func (s *service) GenerateTokensForUser(ctx context.Context, userID int) (Tokens, error) {
	accessToken, err := createAccessToken(userID)
	if err != nil {
		return Tokens{}, err
	}

	refreshToken := uuid.New().String()
	err = s.saveUserToken(ctx, userID, refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *service) saveUserToken(ctx context.Context, userID int, token string) error {
	_, err := s.repository.GetUserToken(ctx, userID)
	if err == ErrTokenForUserNotFound {
		err = s.repository.SaveToken(ctx, token, userID)
	}
	if err != nil {
		return err
	}
	return s.repository.UpdateToken(ctx, token, userID)

}

func createAccessToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(30 * time.Minute),
		"user": userID,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString("secret")
}
