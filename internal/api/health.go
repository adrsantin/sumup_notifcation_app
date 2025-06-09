package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type healthAPIImpl struct{}

func NewHealthAPI(
	r *chi.Mux,
) HealthAPI {
	healthApi := &healthAPIImpl{}

	r.Get("/health", healthApi.Check)

	return healthApi
}

func (h *healthAPIImpl) Check(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check endpoint hit")
	w.Write([]byte("OK"))
}
