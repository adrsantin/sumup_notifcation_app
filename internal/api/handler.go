package api

import (
	"log"
	"sumup/notifications/internal/api/health"

	"github.com/go-chi/chi/v5"
)

func NewAPIs() *chi.Mux {

	log.Println("Initializing API handlers...")
	r := chi.NewRouter()

	log.Println("Setting up health check endpoint...")
	healthAPI := health.NewHealthAPI()

	log.Println("Registering health check route...")
	r.Get("/health", healthAPI.Check)

	return r
}
