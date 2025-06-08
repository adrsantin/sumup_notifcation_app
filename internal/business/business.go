package business

type NotificationService interface {
	SendPaymentNotifications(userID int, amount float64) error
}
