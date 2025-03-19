package errors

import (
	"encoding/json"
	"net/http"
)

// APIError adalah struktur untuk error response
type APIError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Error adalah custom error type
type Error struct {
	Code    int
	Message string
	Err     error
}

// Error mengimplementasikan interface error
func (e *Error) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

// NewError membuat error baru
func NewError(code int, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// WriteError menulis error response ke http.ResponseWriter
func WriteError(w http.ResponseWriter, err error) {
	var apiErr APIError

	// Cek tipe error
	switch e := err.(type) {
	case *Error:
		apiErr = APIError{
			Code:    e.Code,
			Message: e.Message,
			Details: e.Err,
		}
	default:
		apiErr = APIError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Details: err.Error(),
		}
	}

	// Set header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.Code)

	// Encode dan kirim response
	json.NewEncoder(w).Encode(apiErr)
}
