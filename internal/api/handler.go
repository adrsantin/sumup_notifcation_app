package api

import (
	"sumup/notifications/internal/api/health"

	"github.com/go-chi/chi"
)

func NewAPIs() *chi.Mux {
	r := chi.NewRouter()
	healthAPI := health.NewHealthAPI()

	r.Get("/health", healthAPI.Check)

	return r
}
