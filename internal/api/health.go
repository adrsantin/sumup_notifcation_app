package api

import (
	"log"
	"net/http"
)

type healthAPIImpl struct{}

func NewHealthAPI() HealthAPI {
	return &healthAPIImpl{}
}

func (h *healthAPIImpl) Check(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check endpoint hit")
	w.Write([]byte("OK"))
}
