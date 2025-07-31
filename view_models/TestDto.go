package viewmodels

type TestDto struct {
	Name        string            `json:"name"`
	APIEndpoint string            `json:"api_endpoint"`
	Parameters  string            `json:"parameters"`
	Headers     string            `json:"headers"`
	Body        map[string]string `json:"body"`
}
