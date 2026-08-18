package implementations

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/alonso804/ms-go/internal/domain"
)

var (
	ErrStatisticsAPI    = fmt.Errorf("statistics API error")
	ErrStatisticsReq    = fmt.Errorf("statistics request error")
	ErrStatisticsResp   = fmt.Errorf("statistics response error")
	ErrStatisticsDecode = fmt.Errorf("statistics decode error")
)

type APIProvider struct {
	baseURL    string
	httpClient *http.Client
	logger     *slog.Logger
}

func NewAPIProvider(
	baseURL string,
	httpClient *http.Client,
	logger *slog.Logger,
) *APIProvider {
	return &APIProvider{
		baseURL:    baseURL,
		httpClient: httpClient,
		logger:     logger,
	}
}

type statisticsResponse struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Mean       float64 `json:"mean"`
	Sum        float64 `json:"sum"`
	IsDiagonal bool    `json:"isDiagonal"`
}

func (p *APIProvider) GetStatistics(
	ctx context.Context,
	m *domain.Matrix,
) (domain.Statistics, error) {
	p.logger.Info("GetStatistics call")

	payload, err := json.Marshal(map[string]any{
		"matrix": m.Data(),
	})
	if err != nil {
		return domain.Statistics{}, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		p.baseURL+"/get-statistics",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return domain.Statistics{}, ErrStatisticsReq
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		p.logger.Error(
			"statistics API request failed",
			"error", err,
			"url", req.URL.String(),
		)
		return domain.Statistics{}, ErrStatisticsAPI
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return domain.Statistics{}, ErrStatisticsResp
	}

	var result statisticsResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return domain.Statistics{}, ErrStatisticsDecode
	}

	p.logger.Debug("GetStatistics response", "result", result)

	return domain.Statistics{
		Max:        result.Max,
		Min:        result.Min,
		Mean:       result.Mean,
		Sum:        result.Sum,
		IsDiagonal: result.IsDiagonal,
	}, nil
}
