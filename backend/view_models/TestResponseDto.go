package viewmodels

import (
	"time"

	"gorm.io/datatypes"
)

type TestResponseDto struct {
	Name        string         `json:"name"`
	APIEndpoint string         `json:"api_endpoint"`
	Response    datatypes.JSON `json:"response"`
	StatusCode  int            `json:"status_code"`
	CreatedAt   time.Time      `json:"created_at"`
}
