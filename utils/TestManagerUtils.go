package utils

import (
	"apitester/models"
	viewmodels "apitester/view_models"
)

func FormattedResponse(test models.Test) viewmodels.TestResponseDto {
	return viewmodels.TestResponseDto{
		Name:        test.Name,
		APIEndpoint: test.APIEndpoint,
		Response:    test.Response,
		StatusCode:  test.StatusCode,
		CreatedAt:   test.CreatedAt,
	}
}
