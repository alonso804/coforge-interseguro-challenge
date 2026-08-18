package domain

import "errors"

var (
	ErrEmptyMatrix         = errors.New("matrix cannot be empty")
	ErrInvalidMatrix       = errors.New("matrix must be rectangular")
	ErrLinearlyDependent   = errors.New("matrix columns are linearly dependent")
	ErrInvalidQRDimensions = errors.New("invalid dimensions for QR factorization: rows must be greater than or equal to columns")
)
