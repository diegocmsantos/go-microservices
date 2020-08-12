package utils

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
}

// CreateError will create a ApplicationError
func CreateError(message string, statusCode int, code string) *ApplicationError {
	return &ApplicationError{
		Message:    message,
		StatusCode: statusCode,
		Code:       code,
	}
}
