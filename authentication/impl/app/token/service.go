package token

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

var ErrUserWithTokenNotFound = errors.New("user with token not found")

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type User struct {
	Id int
}

type LoginData struct {
	Tokens Tokens
	User   User
}

type Service interface {
	GenerateTokensForUser(ctx context.Context, userID int) (Tokens, error)
	RefreshToken(ctx context.Context, refreshToken string) (Tokens, error)
	RemoveToken(ctx context.Context, token string) error
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

func (s *service) RefreshToken(ctx context.Context, refreshToken string) (Tokens, error) {
	userID, err := s.repository.GetUserByToken(ctx, refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	newAccessToken, err := createAccessToken(userID)
	if err != nil {
		return Tokens{}, err
	}

	newRefreshToken := uuid.New().String()
	err = s.repository.UpdateToken(ctx, newRefreshToken, userID)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *service) RemoveToken(ctx context.Context, token string) error {
	_, err := s.repository.GetUserByToken(ctx, token)
	if err != nil {
		return err
	}
	return s.repository.RemoveToken(ctx, token)
}

func (s *service) saveUserToken(ctx context.Context, userID int, token string) error {
	_, err := s.repository.GetUserToken(ctx, userID)
	if err == ErrUserWithTokenNotFound {
		err = s.repository.SaveToken(ctx, token, userID)
	}
	if err != nil {
		return err
	}
	return s.repository.UpdateToken(ctx, token, userID)

}

func createAccessToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(30 * time.Minute).UnixMilli(),
		"user": userID,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
}
