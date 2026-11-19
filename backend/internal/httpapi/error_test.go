package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type errorBody struct {
	Error struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		RequestID string `json:"requestId"`
	} `json:"error"`
}

func TestErrorEnvelope(t *testing.T) {
	router := NewRouter(Deps{})
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/missing", nil))

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}

	var body errorBody
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if body.Error.Code != "NOT_FOUND" {
		t.Fatalf("code = %q, want NOT_FOUND", body.Error.Code)
	}
	if body.Error.Message == "" {
		t.Fatal("message is empty")
	}
	if body.Error.RequestID == "" {
		t.Fatal("requestId is empty")
	}
	if got := rec.Header().Get(requestIDHeader); got != body.Error.RequestID {
		t.Fatalf("header request id = %q, body = %q", got, body.Error.RequestID)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	router := NewRouter(Deps{})
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/health/live", nil))

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}

	var body errorBody
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if body.Error.Code != "METHOD_NOT_ALLOWED" {
		t.Fatalf("code = %q, want METHOD_NOT_ALLOWED", body.Error.Code)
	}
}

func TestRequestIDSanitized(t *testing.T) {
	router := NewRouter(Deps{})
	req := httptest.NewRequest(http.MethodGet, "/health/live", nil)
	req.Header.Set(requestIDHeader, strings.Repeat("x", 200))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	got := rec.Header().Get(requestIDHeader)
	if got == "" {
		t.Fatal("X-Request-Id header is empty")
	}
	if len(got) > 128 {
		t.Fatalf("generated request id too long: %d", len(got))
	}
}
