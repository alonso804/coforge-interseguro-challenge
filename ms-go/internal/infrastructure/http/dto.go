package http

type RunRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type RunResponse struct {
	RotatedMatrix     [][]float64        `json:"rotatedMatrix"`
	RotatedStatistics StatisticsResponse `json:"rotatedStatistics"`
	Q                 [][]float64        `json:"q"`
	QStatistics       StatisticsResponse `json:"qStatistics"`
	R                 [][]float64        `json:"r"`
	RStatistics       StatisticsResponse `json:"rStatistics"`
}

type StatisticsResponse struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Mean       float64 `json:"mean"`
	Sum        float64 `json:"sum"`
	IsDiagonal bool    `json:"isDiagonal"`
}
