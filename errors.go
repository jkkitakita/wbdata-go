package wbdata

import "fmt"

const (
	ErrInvalidServer = "Invalid Server Error"
)

type ErrorResponse struct {
	Message []ErrorMessage `json:"message"`
}

type ErrorMessage struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%+v", r.Message)
}

type APIError struct {
	Status       int
	ErrorMessage string
}

func (a APIError) Error() string {
	return fmt.Sprintf("%d: %s", a.Status, a.ErrorMessage)
}
