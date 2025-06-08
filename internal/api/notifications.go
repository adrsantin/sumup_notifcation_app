package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sumup/notifications/internal/business"
	"sumup/notifications/internal/entities"
)

type notificationsAPIImpl struct {
	notificationService business.NotificationService
}

func NewNotificationsAPI(
	notificationService business.NotificationService,
) NotificationsAPI {
	return &notificationsAPIImpl{
		notificationService: notificationService,
	}
}

func (n *notificationsAPIImpl) SendPaymentNotifications(w http.ResponseWriter, r *http.Request) {
	log.Println("SendPaymentNotifications endpoint hit")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var notification entities.NotificationDTO
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
