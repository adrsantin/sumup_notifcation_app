package api

import (
	"net/http"
)

type HealthAPI interface {
	Check(w http.ResponseWriter, r *http.Request)
}

type PaymentsAPI interface {
	SendPaymentNotifications(w http.ResponseWriter, r *http.Request)
}
