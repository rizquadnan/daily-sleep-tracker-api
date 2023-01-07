package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/constant"
)

type ErrorResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"statusCode"`
}

func GenerateErrorResponse(message string, statusCode int) ErrorResponse {
	return ErrorResponse{ Message: message, StatusCode: statusCode }
}

func SetStatusNotFoundJSON(c *gin.Context, customMessage string) {
	var message string

	if (customMessage == "") {
		message = constant.GENERIC_NOT_FOUND
	}

	c.JSON(http.StatusNotFound, GenerateErrorResponse(message, http.StatusNotFound))
}

func SetBadRequestJSON(c *gin.Context, customMessage string) {
	var message string

	if (customMessage == "") {
		message = constant.GENERIC_NOT_VALID_PAYLOAD
	}

	c.JSON(http.StatusBadRequest, GenerateErrorResponse(message, http.StatusBadRequest))
}

func SetInternalServerErrorJSON(c *gin.Context, customMessage string) {
	var message string

	if (customMessage == "") {
		message = constant.GENERIC_INTERNAL_SERVER_ERROR
	}

	c.JSON(http.StatusInternalServerError, GenerateErrorResponse(message, http.StatusInternalServerError))
}