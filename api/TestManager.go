package api

import (
	"apitester/models"
	testUtils "apitester/utils"
	viewmodels "apitester/view_models"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func TestManager(e *echo.Group) {
	e.POST("/tests", createTest)
	// e.GET("/tests/:id", getTest)
	// e.PUT("/tests/:id", updateTest)
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
	} else if testDto.APIEndpoint == "http://localhost:8080" {
		return c.JSON(400, map[string]string{"error": "APIEndpoint cannot be this API's endpoint"})
	}

	resp, err := http.Get(testDto.APIEndpoint)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to make API request"})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to read response body"})
	}

	var responseBody any
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return c.JSON(500, map[string]string{
			"error":   "Failed to parse response body",
			"details": err.Error()})
	}

	test := models.Test{
		Name:        testDto.Name,
		APIEndpoint: testDto.APIEndpoint,
		Parameters:  testDto.Parameters,
		Headers:     testDto.Headers,
		Body:        testDto.Body,
		CreatedAt:   time.Now(),
		Response:    responseBody,
		StatusCode:  resp.StatusCode,
	}

	return c.JSON(200, testUtils.FormattedResponse(test))
}
