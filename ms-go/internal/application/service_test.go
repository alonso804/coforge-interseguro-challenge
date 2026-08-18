package application

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/alonso804/ms-go/internal/domain"
)

type stubStatisticsRepository struct {
	stats []domain.Statistics
	err   error
	mu    sync.Mutex
	calls int
}

func (s *stubStatisticsRepository) GetStatistics(_ context.Context, _ *domain.Matrix) (domain.Statistics, error) {
	if s.err != nil {
		return domain.Statistics{}, s.err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.calls >= len(s.stats) {
		return domain.Statistics{}, nil
	}

	stat := s.stats[s.calls]
	s.calls++
	return stat, nil
}

func TestServiceRun_ReturnsAggregatedResults(t *testing.T) {
	repo := &stubStatisticsRepository{
		stats: []domain.Statistics{
			{Max: 9, Min: 1, Mean: 5, Sum: 15, IsDiagonal: false},
			{Max: 6, Min: 2, Mean: 4, Sum: 12, IsDiagonal: false},
			{Max: 7, Min: 3, Mean: 5, Sum: 14, IsDiagonal: false},
		},
	}

	service := NewService(repo)
	result, err := service.Run(context.Background(), [][]float64{{1, 2}, {3, 4}})
	if err != nil {
		t.Fatalf("Run() returned unexpected error: %v", err)
	}

	if repo.calls != 3 {
		t.Fatalf("expected 3 repository calls, got %d", repo.calls)
	}

	if got, want := result.RotatedMatrix.Data(), [][]float64{{3, 1}, {4, 2}}; !equalMatrix(got, want) {
		t.Fatalf("RotatedMatrix = %v, want %v", got, want)
	}
}

func TestServiceRun_PropagatesRepositoryError(t *testing.T) {
	repoErr := errors.New("repository failure")
	repo := &stubStatisticsRepository{err: repoErr}
	service := NewService(repo)

	_, err := service.Run(context.Background(), [][]float64{{1, 2}, {3, 4}})
	if !errors.Is(err, repoErr) {
		t.Fatalf("Run() error = %v, want wrapped %v", err, repoErr)
	}
}

func equalMatrix(a, b [][]float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
