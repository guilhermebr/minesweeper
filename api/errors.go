package api

import (
	"encoding/json"
	"net/http"
)

var (
	ErrInternalServer = Error{StatusCode: http.StatusInternalServerError, Type: "server_error", Message: "Internal server error. The error has been logged and we are working on it"}
	ErrInvalidJSON    = Error{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
	ErrAlreadyExists  = Error{StatusCode: http.StatusConflict, Type: "already_exists", Message: "Another resource has the same value as this field"}
)

type Error struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func (e Error) Send(w http.ResponseWriter) error {
	statusCode := e.StatusCode
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(e)
}
