package handler

import "fmt"

//HTTPError model
type HTTPError struct {
	Message    string
	StatusCode int
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s%d", e.Message, e.StatusCode)
}

//NewHTTPError creation
func NewHTTPError(message string, statusCode int) error {
	return &HTTPError{Message: message, StatusCode: statusCode}
}
