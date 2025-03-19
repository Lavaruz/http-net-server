package errors

import "net/http"

// Predefined errors
var (
	// ErrNotFound adalah error untuk resource yang tidak ditemukan
	ErrNotFound = NewError(http.StatusNotFound, "Resource not found", nil)

	// ErrBadRequest adalah error untuk request yang tidak valid
	ErrBadRequest = NewError(http.StatusBadRequest, "Bad request", nil)

	// ErrUnauthorized adalah error untuk akses yang tidak diizinkan
	ErrUnauthorized = NewError(http.StatusUnauthorized, "Unauthorized", nil)

	// ErrForbidden adalah error untuk akses yang dilarang
	ErrForbidden = NewError(http.StatusForbidden, "Forbidden", nil)

	// ErrInternalServer adalah error untuk masalah server internal
	ErrInternalServer = NewError(http.StatusInternalServerError, "Internal server error", nil)

	// ErrValidation adalah error untuk validasi yang gagal
	ErrValidation = NewError(http.StatusUnprocessableEntity, "Validation failed", nil)

	// ErrDatabase adalah error untuk masalah database
	ErrDatabase = NewError(http.StatusInternalServerError, "Database error", nil)
)
