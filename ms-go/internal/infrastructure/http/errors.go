package http

import (
	"errors"

	"github.com/alonso804/ms-go/internal/domain"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func mapError(err error) error {
	switch {
	case errors.Is(err, domain.ErrEmptyMatrix):
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)

	case errors.Is(err, domain.ErrInvalidMatrix):
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)

	case errors.Is(err, domain.ErrInvalidQRDimensions):
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)

	case errors.Is(err, domain.ErrLinearlyDependent):
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)

	default:
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"internal server error",
		)
	}
}

func parseValidationError(c fiber.Ctx, err error) error {
	if validationErrors, ok := errors.AsType[validator.ValidationErrors](err); ok {
		out := make([]fiber.Map, 0, len(validationErrors))

		for _, e := range validationErrors {
			out = append(out, fiber.Map{
				"field": e.Field(),
				"rule":  e.Tag(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": out})
	}

	return err
}
