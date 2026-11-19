package httpapi

import (
	"context"
	"net/http"
	"time"

	"github.com/LeGiaBaoDev/Go_Wander/backend/internal/db"
)

type healthHandler struct {
	pinger  db.Pinger
	timeout time.Duration
}

func newHealthHandler(pinger db.Pinger) healthHandler {
	return healthHandler{pinger: pinger, timeout: 2 * time.Second}
}

func (h healthHandler) live(w http.ResponseWriter, _ *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h healthHandler) ready(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), h.timeout)
	defer cancel()

	checks := map[string]string{"database": "ok"}
	status := http.StatusOK

	if h.pinger == nil {
		checks["database"] = "not_configured"
		status = http.StatusServiceUnavailable
	} else if err := h.pinger.Ping(ctx); err != nil {
		checks["database"] = "unavailable"
		status = http.StatusServiceUnavailable
	}

	body := map[string]any{"status": "ok", "checks": checks}
	if status != http.StatusOK {
		body["status"] = "unavailable"
	}
	WriteJSON(w, status, body)
}
