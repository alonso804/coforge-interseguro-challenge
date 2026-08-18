package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v3"

	"github.com/alonso804/ms-go/internal/infrastructure/implementations"
)

func Auth(provider *implementations.JWTProvider) fiber.Handler {
	return func(c fiber.Ctx) error {
		header := c.Get("Authorization")

		if header == "" {
			return fiber.ErrUnauthorized
		}

		const prefix = "Bearer "

		if !strings.HasPrefix(header, prefix) {
			return fiber.ErrUnauthorized
		}

		token := strings.TrimSpace(strings.TrimPrefix(header, prefix))

		if token == "" {
			return fiber.ErrUnauthorized
		}

		if err := provider.Validate(token); err != nil {
			return fiber.ErrUnauthorized
		}

		return c.Next()
	}
}
