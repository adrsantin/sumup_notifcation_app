package payments

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sumup/notifications/internal/api"
	"sumup/notifications/internal/business"
	"sumup/notifications/internal/entities"

	"github.com/go-chi/chi/v5"
)

type paymentsAPIImpl struct {
	paymentsService business.PaymentService
}

func NewPaymentsAPI(
	r *chi.Mux,
	paymentsService business.PaymentService,
) api.PaymentsAPI {
	paymentsApi := &paymentsAPIImpl{
		paymentsService: paymentsService,
	}

	r.Post("/payments/notification", paymentsApi.SendPaymentNotifications)

	return paymentsApi
}

func (n *paymentsAPIImpl) SendPaymentNotifications(w http.ResponseWriter, r *http.Request) {
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
	err = n.paymentsService.ProcessPaymentNotification(notification.UserID, notification.Amount)
	if err != nil {
		http.Error(w, "Failed to send notifications: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
