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

func TestManager(e *echo.Group) {
	e.POST("/tests", createTest)
	e.GET("/tests/:id", getTest)
	e.PUT("/tests/:id", updateTest)
	// e.DELETE("/tests/:id", deleteTest)
	// e.GET("/tests", listTests)
}

func createTest(c echo.Context) error {
	var testDto viewmodels.TestDto
	if err := c.Bind(&testDto); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}
	if testDto.Name == "" || testDto.APIEndpoint == "" {
		return c.JSON(400, map[string]string{"error": "Name and APIEndpoint are required"})
	}

	if err := testUtils.ValidateAPIEndpoint(testDto.APIEndpoint); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	client := &http.Client{Timeout: 10 * time.Second}

	method := strings.ToUpper(testDto.Method)
	var rbody io.Reader

	if testDto.Body != nil {
		rbody = bytes.NewReader(testDto.Body)
	}

	if !validMethods[method] {
		return c.JSON(400, map[string]string{"error": "Invalid HTTP method"})
	}

	if method == http.MethodGet {
		rbody = nil // GET requests should not have a body
	}

	req, err := http.NewRequest(method, testDto.APIEndpoint, rbody)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
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
		return c.JSON(500, map[string]string{"error": "Failed to read response body"})
	}

	if !json.Valid(body) {
		return c.JSON(400, map[string]string{"error": "Response is not valid JSON"})
	}
	responseBody := datatypes.JSON(body)

	parameters, err := json.Marshal(testDto.Parameters)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to marshal parameters"})
	}

	headers, err := json.Marshal(testDto.Headers)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to marshal headers"})
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

	db := database.GetDB()
	if err := db.Create(&test).Error; err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to save test to database"})
	}

	return c.JSON(http.StatusCreated, testUtils.FormattedResponse(test))
}

func getTest(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(400, map[string]string{"error": "Test ID is required"})
	}

	id_int, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid Test ID"})
	}

	test, err := database.GetTestByID(id_int)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Test not found"})
	}

	return c.JSON(200, testUtils.FormattedResponse(*test))

}

func updateTest(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(400, map[string]string{"error": "Test ID is required"})
	}

	var testDto viewmodels.TestDto
	if err := c.Bind(&testDto); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	if testDto.Name == "" || testDto.APIEndpoint == "" {
		return c.JSON(400, map[string]string{"error": "Name and APIEndpoint are required"})
	} else if testDto.APIEndpoint == "http://localhost:8080" {
		return c.JSON(400, map[string]string{"error": "APIEndpoint cannot be this API's endpoint"})
	}

	id_int, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid Test ID"})
	}

	existingTest, err := database.GetTestByID(id_int)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Test not found"})
	}

	parameters, err := json.Marshal(testDto.Parameters)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to marshal parameters"})
	}

	headers, err := json.Marshal(testDto.Headers)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to marshal headers"})
	}

	existingTest.Name = testDto.Name
	existingTest.APIEndpoint = testDto.APIEndpoint
	existingTest.Parameters = parameters
	existingTest.Headers = headers
	existingTest.Body = testDto.Body

	err = database.UpdateTest(existingTest)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to update test"})
	}

	return c.JSON(200, testUtils.FormattedResponse(*existingTest))
}

var validMethods = map[string]bool{
	http.MethodGet:    true,
	http.MethodPost:   true,
	http.MethodPut:    true,
	http.MethodDelete: true,
}
