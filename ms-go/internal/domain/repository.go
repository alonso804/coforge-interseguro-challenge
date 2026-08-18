package domain

type Repository interface {
	Rotate(matrix [][]float64) [][]float64

	GetQrFactorization(matrix [][]float64) ([][]float64, [][]float64)

	GetStatistics(matrix [][]float64) (max float64, min float64, mean float64, sum float64, isDiagonal bool)
}
