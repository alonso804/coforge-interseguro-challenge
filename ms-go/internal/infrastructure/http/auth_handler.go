package http

import (
	"errors"

	"github.com/gofiber/fiber/v3"

	"github.com/alonso804/ms-go/internal/application"
)

type AuthHandler struct {
	service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (ah *AuthHandler) Run(c fiber.Ctx) error {
	var request LoginRequest

	if err := c.Bind().Body(&request); err != nil {
		return parseValidationError(c, err)
	}

	token, err := ah.service.Login(
		c.Context(),
		request.Username,
		request.Password,
	)
	if err != nil {
		if errors.Is(err, application.ErrInvalidCredentials) {
			return fiber.ErrUnauthorized
		}

		return err
	}

	return c.JSON(LoginResponse{
		Token: token,
	})
}
