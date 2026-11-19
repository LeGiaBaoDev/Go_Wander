package httpapi

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type errorEnvelope struct {
	Error errorDetail `json:"error"`
}

type errorDetail struct {
	Code      string         `json:"code"`
	Message   string         `json:"message"`
	RequestID string         `json:"requestId"`
	Details   map[string]any `json:"details,omitempty"`
}

// WriteJSON writes a JSON response with the shared content type.
func WriteJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		slog.Error("write json response", "error", err)
	}
}

// WriteError writes the standard error envelope from the API contract.
func WriteError(w http.ResponseWriter, r *http.Request, status int, code, message string, details map[string]any) {
	WriteJSON(w, status, errorEnvelope{Error: errorDetail{
		Code:      code,
		Message:   message,
		RequestID: RequestIDFrom(r.Context()),
		Details:   details,
	}})
}
