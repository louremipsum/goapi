// Package api provides the core API structures and error handling functionality
// for the coin balance checking service.
package api

import (
	"encoding/json"
	"net/http"
)

// CoinBalanceParams defines the expected parameters for coin balance requests
type CoinBalanceParams struct {
	Username string
}

// CoinBalanceResponse represents the API response for coin balance queries
type CoinBalanceResponse struct {
	Code    int   // HTTP status code
	Balance int64 // User's coin balance
}

// Error defines the structure for API error responses
type Error struct {
	Code    int    // HTTP status code
	Message string // Error description
}

// writeError is a helper function to send error responses in JSON format
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

// Pre-defined error handlers for common API responses
var (
	// RequestErrorHandler handles bad request errors (400)
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	// InternalErrorHandler handles internal server errors (500)
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error happened.", http.StatusInternalServerError)
	}
)
