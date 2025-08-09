package models

import (
	"time"

	"gorm.io/datatypes"
)

type Test struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	APIEndpoint string         `json:"api_endpoint"`
	Parameters  datatypes.JSON `json:"parameters"`
	Headers     datatypes.JSON `json:"headers"`
	Body        datatypes.JSON `json:"body"`
	CreatedAt   time.Time      `json:"created_at"`
	Response    datatypes.JSON `json:"response"`
	StatusCode  int            `json:"status_code"`
}
