package viewmodels

import "gorm.io/datatypes"

type TestDto struct {
	Name        string            `json:"name"`
	APIEndpoint string            `json:"api_endpoint"`
	Method      string            `json:"method"`
	Parameters  map[string]string `json:"parameters"`
	Headers     map[string]string `json:"headers"`
	Body        datatypes.JSON    `json:"body"`
}
