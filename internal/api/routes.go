package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(
	r *chi.Mux,
	healthAPI HealthAPI,
	notificationsAPI NotificationsAPI,
) {
	r.Get("/health", healthAPI.Check)
	r.Post("/notifications/send", notificationsAPI.SendPaymentNotifications)
}

type HealthAPI interface {
	Check(w http.ResponseWriter, r *http.Request)
}

type NotificationsAPI interface {
	SendPaymentNotifications(w http.ResponseWriter, r *http.Request)
}
