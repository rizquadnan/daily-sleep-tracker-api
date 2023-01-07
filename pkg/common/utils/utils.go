package utils

type ErrorResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"statusCode"`
}

func GenerateErrorResponse(message string, statusCode int) ErrorResponse {
	return ErrorResponse{ Message: message, StatusCode: statusCode }
}