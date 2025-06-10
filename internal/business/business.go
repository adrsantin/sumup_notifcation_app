package business

import "sumup/notifications/internal/entities"

//go:generate mockgen -source=business.go -destination=../mocks/business.go -package=mocks
type PaymentService interface {
	ProcessPaymentNotification(userID int, amount float64) error
}

type NotificationService interface {
	SendNotification(user *entities.User, amount float64, notificationType entities.NotificationType) error
}
