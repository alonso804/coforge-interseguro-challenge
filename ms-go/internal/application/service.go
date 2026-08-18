// Package application layer is responsible for orchestrating the domain logic and coordinating the flow of data between the domain layer and the infrastructure layer. It contains the business logic of the application and defines the use cases that the application supports.
package application

import (
	"context"

	"github.com/alonso804/ms-go/internal/domain"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	statisticsRepository domain.StatisticsRepositoy
}

func NewService(
	statisticsRepository domain.StatisticsRepositoy,
) *Service {
	return &Service{
		statisticsRepository: statisticsRepository,
	}
}

type Result struct {
	RotatedMatrix     *domain.Matrix
	RotatedStatistics domain.Statistics

	Q           *domain.Matrix
	QStatistics domain.Statistics

	R           *domain.Matrix
	RStatistics domain.Statistics
}

func (s *Service) Run(
	ctx context.Context,
	matrix [][]float64,
) (*Result, error) {
	m, err := domain.NewMatrix(matrix)
	if err != nil {
		return nil, err
	}

	m.Rotate()

	q, r, err := m.QR()
	if err != nil {
		return nil, err
	}

	var (
		rotatedStatistics domain.Statistics
		qStatistics       domain.Statistics
		rStatistics       domain.Statistics
	)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		rotatedStatistics, err = s.statisticsRepository.GetStatistics(ctx, m)

		return err
	})

	g.Go(func() error {
		qStatistics, err = s.statisticsRepository.GetStatistics(ctx, q)

		return err
	})

	g.Go(func() error {
		rStatistics, err = s.statisticsRepository.GetStatistics(ctx, r)

		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &Result{
		RotatedMatrix:     m,
		RotatedStatistics: rotatedStatistics,

		Q:           q,
		QStatistics: qStatistics,

		R:           r,
		RStatistics: rStatistics,
	}, nil
}
