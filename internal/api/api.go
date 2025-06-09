package api

import (
	"net/http"
)

type HealthAPI interface {
	Check(w http.ResponseWriter, r *http.Request)
}

type NotificationsAPI interface {
	SendPaymentNotifications(w http.ResponseWriter, r *http.Request)
}
