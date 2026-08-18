package application

import (
	"context"
	"errors"

	"github.com/alonso804/ms-go/internal/infrastructure/implementations"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	jwtProvider *implementations.JWTProvider
}

func NewAuthService(jwtProvider *implementations.JWTProvider) *AuthService {
	return &AuthService{
		jwtProvider,
	}
}

const (
	USERNAME = "admin"
	PASSWORD = "password"
)

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	if username != USERNAME || password != PASSWORD {
		return "", ErrInvalidCredentials
	}

	token, err := s.jwtProvider.Generate(username)
	if err != nil {
		return "", err
	}

	return token, nil
}
