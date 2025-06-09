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

type RequestDTO struct {
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
}

type MessageDTO struct {
	UserID           int              `json:"user_id"`
	Name             string           `json:"name"`
	Email            string           `json:"email"`
	Phone            string           `json:"phone"`
	NotificationType NotificationType `json:"notification_type"`
}
