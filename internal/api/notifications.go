package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sumup/notifications/internal/business"
	"sumup/notifications/internal/entities"

	"github.com/go-chi/chi/v5"
)

type notificationsAPIImpl struct {
	notificationService business.NotificationService
}

func NewNotificationsAPI(
	r *chi.Mux,
	notificationService business.NotificationService,
) NotificationsAPI {
	notificationsApi := &notificationsAPIImpl{
		notificationService: notificationService,
	}

	r.Post("/notifications/send", notificationsApi.SendPaymentNotifications)

	return notificationsApi
}

func (n *notificationsAPIImpl) SendPaymentNotifications(w http.ResponseWriter, r *http.Request) {
	log.Println("SendPaymentNotifications endpoint hit")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var notification entities.RequestDTO
	err = json.Unmarshal(body, &notification)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = n.notificationService.SendPaymentNotifications(notification.UserID, notification.Amount)
	if err != nil {
		http.Error(w, "Failed to send notifications: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
