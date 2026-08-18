package domain

import "context"

type Statistics struct {
	Max        float64
	Min        float64
	Mean       float64
	Sum        float64
	IsDiagonal bool
}

type StatisticsRepositoy interface {
	GetStatistics(ctx context.Context, matrix *Matrix) (Statistics, error)
}
