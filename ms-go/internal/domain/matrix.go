// Package domain contains the Matrix struct and its associated methods for matrix operations, including QR decomposition.
package domain

import (
	"math"
)

const linearDependenceTolerance = 1e-12

type Matrix struct {
	data [][]float64
	rows int
	cols int
}

func NewMatrix(data [][]float64) (*Matrix, error) {
	if len(data) == 0 || len(data[0]) == 0 {
		return nil, ErrEmptyMatrix
	}

	cols := len(data[0])

	for _, row := range data {
		if len(row) != cols {
			return nil, ErrInvalidMatrix
		}
	}

	copied := make([][]float64, len(data))

	for i := range data {
		copied[i] = make([]float64, cols)
		copy(copied[i], data[i])
	}

	return &Matrix{
		data: copied,
		rows: len(data),
		cols: cols,
	}, nil
}

func NewEmptyMatrix(rows, cols int) *Matrix {
	data := make([][]float64, rows)

	for i := range data {
		data[i] = make([]float64, cols)
	}

	return &Matrix{
		data: data,
		rows: rows,
		cols: cols,
	}
}

func (m *Matrix) Data() [][]float64 {
	data := make([][]float64, m.rows)

	for i := range m.data {
		data[i] = make([]float64, m.cols)
		copy(data[i], m.data[i])
	}

	return data
}

func (m *Matrix) Rotate() {
	newMatrix := NewEmptyMatrix(m.cols, m.rows)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			newMatrix.data[j][m.rows-1-i] = m.data[i][j]
		}
	}

	m.data = newMatrix.data
	m.rows, m.cols = newMatrix.rows, newMatrix.cols
}

func (m *Matrix) QR() (*Matrix, *Matrix, error) {
	if m.rows == 0 || m.cols == 0 {
		return nil, nil, ErrEmptyMatrix
	}

	if m.rows < m.cols {
		return nil, nil, ErrInvalidQRDimensions
	}

	q := NewEmptyMatrix(m.rows, m.cols)
	r := NewEmptyMatrix(m.cols, m.cols)

	v := NewEmptyMatrix(m.rows, m.cols)

	for i := 0; i < m.rows; i++ {
		copy(v.data[i], m.data[i])
	}

	for j := 0; j < m.cols; j++ {
		norm := v.columnNorm(j)

		if norm < linearDependenceTolerance {
			return nil, nil, ErrLinearlyDependent
		}

		r.data[j][j] = norm

		for i := 0; i < m.rows; i++ {
			q.data[i][j] = v.data[i][j] / norm
		}

		for k := j + 1; k < m.cols; k++ {
			r.data[j][k] = q.dotColumn(v, j, k)

			for i := 0; i < m.rows; i++ {
				v.data[i][k] -= r.data[j][k] * q.data[i][j]
			}
		}
	}

	return q, r, nil
}

func (m *Matrix) columnNorm(column int) float64 {
	var sum float64

	for i := 0; i < m.rows; i++ {
		sum += m.data[i][column] * m.data[i][column]
	}

	return math.Sqrt(sum)
}

func (m *Matrix) dotColumn(matrix *Matrix, qColumn, mColumn int) float64 {
	var result float64

	for i := 0; i < m.rows; i++ {
		result += m.data[i][qColumn] * matrix.data[i][mColumn]
	}

	return result
}
