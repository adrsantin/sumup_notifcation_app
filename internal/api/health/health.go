package health

import (
	"log"
	"net/http"
	"sumup/notifications/internal/api"

	"github.com/go-chi/chi/v5"
)

type healthAPIImpl struct{}

func NewHealthAPI(
	r *chi.Mux,
) api.HealthAPI {
	healthApi := &healthAPIImpl{}

	r.Get("/health", healthApi.Check)

	return healthApi
}

func (h *healthAPIImpl) Check(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check endpoint hit")
	w.Write([]byte("OK"))
}
