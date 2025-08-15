package api

import (
	database "apitester/database"
	"apitester/models"
	testUtils "apitester/utils"
	viewmodels "apitester/view_models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

func TestManager(e *echo.Group, repo *database.TestRepository) {
	h := &Handler{repo: repo}

	e.POST("/tests", h.createTest)
	e.GET("/tests/:id", h.getTest)
	e.GET("/tests", h.getAllTests)
	e.PUT("/tests/:id", h.updateTest)
	e.DELETE("/tests/:id", h.deleteTest)
}

type Handler struct {
	repo *database.TestRepository
}

func (h *Handler) getAllTests(c echo.Context) error {
	tests, err := h.repo.GetAllTests()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve tests"})
	}

	var testResponses = make([]viewmodels.TestResponseDto, len(tests))
	for i, test := range tests {
		testResponses[i] = testUtils.FormattedResponse(test)
	}

	return c.JSON(http.StatusOK, testResponses)
}

func (h *Handler) createTest(c echo.Context) error {
	var testDto viewmodels.TestDto
	if err := c.Bind(&testDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if testDto.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name is required"})
	}

	if err := testUtils.ValidateAPIEndpoint(testDto.APIEndpoint); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	client := &http.Client{Timeout: 10 * time.Second}

	method := strings.ToUpper(testDto.Method)
	var rbody io.Reader

	if testDto.Body != nil {
		rbody = bytes.NewReader(testDto.Body)
	}

	if !validMethods[method] {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid HTTP method"})
	}

	if method == http.MethodGet {
		rbody = nil // GET requests should not have a body
	}

	req, err := http.NewRequest(method, testDto.APIEndpoint, rbody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	for key, value := range testDto.Headers {
		req.Header.Add(key, value)
	}

	q := req.URL.Query()
	for key, value := range testDto.Parameters {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(502, map[string]string{"error": "Failed to make API request"})
	}
	defer resp.Body.Close()

	const max = 1 << 20 // 1MB
	body, err := io.ReadAll(io.LimitReader(resp.Body, max))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read response body"})
	}

	if !json.Valid(body) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Response is not valid JSON"})
	}
	responseBody := datatypes.JSON(body)

	parameters, err := json.Marshal(testDto.Parameters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal parameters"})
	}

	headers, err := json.Marshal(testDto.Headers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal headers"})
	}

	test := models.Test{
		Name:        testDto.Name,
		APIEndpoint: testDto.APIEndpoint,
		Parameters:  parameters,
		Headers:     headers,
		Body:        testDto.Body,
		Response:    responseBody,
		StatusCode:  resp.StatusCode,
	}

	if err = h.repo.CreateTest(&test); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save test to database"})
	}

	return c.JSON(http.StatusCreated, testUtils.FormattedResponse(test))
}

func (h *Handler) getTest(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Test ID is required"})
	}

	id_int, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Test ID"})
	}

	test, err := h.repo.GetTestByID(id_int)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Test not found"})
	}

	return c.JSON(http.StatusOK, testUtils.FormattedResponse(*test))

}

func (h *Handler) updateTest(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Test ID is required"})
	}

	var testDto viewmodels.TestDto
	if err := c.Bind(&testDto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if testDto.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name is required"})
	}

	if err := testUtils.ValidateAPIEndpoint(testDto.APIEndpoint); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	id_int, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Test ID"})
	}

	existingTest, err := h.repo.GetTestByID(id_int)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Test not found"})
	}

	parameters, err := json.Marshal(testDto.Parameters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal parameters"})
	}

	headers, err := json.Marshal(testDto.Headers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal headers"})
	}

	existingTest.Name = testDto.Name
	existingTest.APIEndpoint = testDto.APIEndpoint
	existingTest.Parameters = parameters
	existingTest.Headers = headers
	existingTest.Body = testDto.Body

	err = h.repo.UpdateTest(existingTest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update test"})
	}

	return c.JSON(http.StatusOK, testUtils.FormattedResponse(*existingTest))
}

func (h *Handler) deleteTest(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Test ID is required"})
	}

	id_int, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Test ID"})
	}

	if err = h.repo.DeleteTest(id_int); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Fail while deleting test"})
	}

	return c.JSON(http.StatusOK, map[string]string{"result": "Test was deleted successfully"})
}

var validMethods = map[string]bool{
	http.MethodGet:    true,
	http.MethodPost:   true,
	http.MethodPut:    true,
	http.MethodDelete: true,
}
