package entities

type NotificationType string

const (
	EmailNotification NotificationType = "email"
	SMSNotification   NotificationType = "sms"
)

var validNotificationTypes = map[NotificationType]bool{
	EmailNotification: true,
	SMSNotification:   true,
}

func IsValidNotificationType(s string) bool {
	_, ok := validNotificationTypes[NotificationType(s)]
	return ok
}

type NotificationDTO struct {
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
}
