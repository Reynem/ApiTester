package models

import (
	"time"
)

type Test struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	APIEndpoint string            `json:"api_endpoint"`
	Parameters  string            `json:"parameters"`
	Headers     string            `json:"headers"`
	Body        map[string]string `json:"body"`
	CreatedAt   time.Time         `json:"created_at"`
	Response    any               `json:"response"`
	StatusCode  int               `json:"status_code"`
}
