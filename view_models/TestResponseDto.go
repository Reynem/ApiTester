package viewmodels

import (
	"time"
)

type TestResponseDto struct {
	Name        string    `json:"name"`
	APIEndpoint string    `json:"api_endpoint"`
	Response    any       `json:"response"`
	StatusCode  int       `json:"status_code"`
	CreatedAt   time.Time `json:"created_at"`
}
