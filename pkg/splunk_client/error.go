package splunk_client

import (
	"fmt"
	"net/http"
	"strings"
)

// APIError represents custom API errors.
type APIError struct {
	StatusCode int
	Message    string
}

// Error returns the error message.
func (e *APIError) Error() string {
	return fmt.Sprintf("API Error: %d - %s", e.StatusCode, e.Message)
}

func (r *APIError) Notfound() bool {
	return r.StatusCode == http.StatusNotFound // 403
}

func (r *APIError) AlreadyExist() bool {
	return r.StatusCode == http.StatusUnprocessableEntity && strings.Contains(r.Message, "already exists") // 422
}
