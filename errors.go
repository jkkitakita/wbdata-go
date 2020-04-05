package wbdata

import (
	"fmt"
)

const (
	// ErrInvalidServer is an error message for "Internal Server Error"
	ErrInvalidServer = "Internal Server Error"
)

type (
	// ErrorResponse is a struct for error response
	ErrorResponse struct {
		URL     string
		Code    int
		Message []ErrorMessage `json:"message"`
	}

	// ErrorMessage is a struct for error message
	ErrorMessage struct {
		ID    string `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	// APIError is a struct for API's error
	APIError struct {
		URL     string
		Code    int
		Message string
	}
)

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("msg: %s, code: %d, URL: %s", e.Message, e.Code, e.URL)
}

func (ae *APIError) Error() string {
	return fmt.Sprintf("%s returned status code %d: %s", ae.URL, ae.Code, ae.Message)
}

// NewAPIError returns an APIError struct
func NewAPIError(url string, code int, msg string) *APIError {
	return &APIError{
		URL:     url,
		Code:    code,
		Message: msg,
	}
}
