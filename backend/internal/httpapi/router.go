package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/LeGiaBaoDev/Go_Wander/backend/internal/db"

	"github.com/go-chi/chi/v5"
)

// Deps are the external dependencies of the HTTP API.
type Deps struct {
	Logger *slog.Logger
	DB     db.Pinger
}

// NewRouter builds the HTTP router for the R1 scaffold.
func NewRouter(deps Deps) http.Handler {
	if deps.Logger == nil {
		deps.Logger = slog.Default()
	}
	health := newHealthHandler(deps.DB)

	router := chi.NewRouter()
	router.Use(RequestID)
	router.Use(RequestLogger(deps.Logger))
	router.Use(Recoverer(deps.Logger))

	router.Get("/health/live", health.live)
	router.Get("/health/ready", health.ready)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		WriteError(w, r, http.StatusNotFound, "NOT_FOUND", "Không tìm thấy tài nguyên.", nil)
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		WriteError(w, r, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Phương thức không được hỗ trợ.", nil)
	})

	return router
}
