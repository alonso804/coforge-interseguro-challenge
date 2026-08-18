// Package http provides the HTTP handler for the application.
package http

import (
	"github.com/alonso804/ms-go/internal/domain"
	"github.com/gofiber/fiber/v3"

	"github.com/alonso804/ms-go/internal/application"
)

type Handler struct {
	service *application.Service
}

func NewHandler(service *application.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Run(c fiber.Ctx) error {
	var request RunRequest

	if err := c.Bind().JSON(&request); err != nil {
		return parseValidationError(c, err)
	}

	result, err := h.service.Run(
		c.Context(),
		request.Matrix,
	)
	if err != nil {
		return mapError(err)
	}

	return c.JSON(RunResponse{
		RotatedMatrix:     result.RotatedMatrix.Data(),
		RotatedStatistics: toStatisticsResponse(result.RotatedStatistics),
		Q:                 result.Q.Data(),
		QStatistics:       toStatisticsResponse(result.QStatistics),
		R:                 result.R.Data(),
		RStatistics:       toStatisticsResponse(result.RStatistics),
	})
}

func toStatisticsResponse(
	statistics domain.Statistics,
) StatisticsResponse {
	return StatisticsResponse{
		Max:        statistics.Max,
		Min:        statistics.Min,
		Mean:       statistics.Mean,
		Sum:        statistics.Sum,
		IsDiagonal: statistics.IsDiagonal,
	}
}
