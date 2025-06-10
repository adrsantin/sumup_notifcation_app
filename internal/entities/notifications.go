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

type MessageDTO struct {
	UserID           int              `json:"user_id"`
	Name             string           `json:"name"`
	Email            string           `json:"email"`
	Phone            string           `json:"phone"`
	Amount           float64          `json:"amount"`
	NotificationType NotificationType `json:"notification_type"`
}
