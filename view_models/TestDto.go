package viewmodels

import "gorm.io/datatypes"

type TestDto struct {
	Name        string         `json:"name"`
	APIEndpoint string         `json:"api_endpoint"`
	Parameters  string         `json:"parameters"`
	Headers     datatypes.JSON `json:"headers"`
	Body        datatypes.JSON `json:"body"`
}
