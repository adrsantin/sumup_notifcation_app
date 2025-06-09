package business

//go:generate mockgen -source=business.go -destination=../mocks/business.go -package=mocks
type NotificationService interface {
	SendPaymentNotifications(userID int, amount float64) error
}
