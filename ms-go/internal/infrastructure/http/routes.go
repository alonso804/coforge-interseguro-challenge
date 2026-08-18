package http

import (
	"github.com/alonso804/ms-go/internal/infrastructure/http/middlewares"
	"github.com/alonso804/ms-go/internal/infrastructure/implementations"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(
	app *fiber.App,
	handler *Handler,
	authHandler *AuthHandler,
	jwtProvider *implementations.JWTProvider,
) {
	app.Post("/login", authHandler.Run)

	api := app.Group("/")

	api.Use(middlewares.Auth(jwtProvider))

	api.Post("/matrix", handler.Run)
}
