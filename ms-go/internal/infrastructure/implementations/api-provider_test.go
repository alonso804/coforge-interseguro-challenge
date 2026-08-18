package implementations

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alonso804/ms-go/internal/domain"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestAPIProviderGetStatistics_Success(t *testing.T) {
	var payload struct {
		Matrix [][]float64 `json:"matrix"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/get-statistics" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Content-Type"); !strings.Contains(got, "application/json") {
			t.Fatalf("unexpected content type: %q", got)
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("decode request payload: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"max":9,"min":1,"mean":5,"sum":15,"isDiagonal":false}`)
	}))
	defer server.Close()

	provider := NewAPIProvider(server.URL, server.Client(), slog.New(slog.NewTextHandler(io.Discard, nil)))
	matrix, err := domain.NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	if err != nil {
		t.Fatalf("NewMatrix() unexpected error: %v", err)
	}

	stats, err := provider.GetStatistics(context.Background(), matrix)
	if err != nil {
		t.Fatalf("GetStatistics() returned unexpected error: %v", err)
	}

	want := domain.Statistics{Max: 9, Min: 1, Mean: 5, Sum: 15, IsDiagonal: false}
	if stats != want {
		t.Fatalf("GetStatistics() = %#v, want %#v", stats, want)
	}

	if got, want := payload.Matrix, [][]float64{{1, 2, 3}, {4, 5, 6}}; !equalMatrix(got, want) {
		t.Fatalf("request payload matrix = %v, want %v", got, want)
	}
}

func TestAPIProviderGetStatistics_RequestFailure(t *testing.T) {
	provider := NewAPIProvider("http://example.com", &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("transport failed")
	})}, slog.New(slog.NewTextHandler(io.Discard, nil)))

	matrix, err := domain.NewMatrix([][]float64{{1, 2}, {3, 4}})
	if err != nil {
		t.Fatalf("NewMatrix() unexpected error: %v", err)
	}

	_, err = provider.GetStatistics(context.Background(), matrix)
	if !errors.Is(err, ErrStatisticsAPI) {
		t.Fatalf("GetStatistics() error = %v, want %v", err, ErrStatisticsAPI)
	}
}

func TestAPIProviderGetStatistics_NonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request", http.StatusBadRequest)
	}))
	defer server.Close()

	provider := NewAPIProvider(server.URL, server.Client(), slog.New(slog.NewTextHandler(io.Discard, nil)))
	matrix, err := domain.NewMatrix([][]float64{{1, 2}, {3, 4}})
	if err != nil {
		t.Fatalf("NewMatrix() unexpected error: %v", err)
	}

	_, err = provider.GetStatistics(context.Background(), matrix)
	if !errors.Is(err, ErrStatisticsResp) {
		t.Fatalf("GetStatistics() error = %v, want %v", err, ErrStatisticsResp)
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
