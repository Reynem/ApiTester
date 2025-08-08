package api

import (
	database "apitester/database"
	"apitester/models"
	testUtils "apitester/utils"
	viewmodels "apitester/view_models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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
	if testDto.APIEndpoint == "http://localhost:8080" {
		return c.JSON(400, map[string]string{"error": "APIEndpoint cannot be this API's endpoint"})
	}

	// TODO: Add normal URL validation (schemes, private IP, loopback)
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, testDto.APIEndpoint, nil)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid APIEndpoint"})
	}

	// TODO: Add headers and parameters to the request
	// for k, v := range testDto.Headers { req.Header.Set(k, v) }

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

	test := models.Test{
		Name:        testDto.Name,
		APIEndpoint: testDto.APIEndpoint,
		Parameters:  testDto.Parameters,
		Headers:     testDto.Headers,
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

	existingTest.Name = testDto.Name
	existingTest.APIEndpoint = testDto.APIEndpoint
	existingTest.Parameters = testDto.Parameters
	existingTest.Headers = testDto.Headers
	existingTest.Body = testDto.Body

	err = database.UpdateTest(existingTest)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to update test"})
	}

	return c.JSON(200, testUtils.FormattedResponse(*existingTest))
}
