package users

type NotificationType string

const (
	EmailNotification NotificationType = "email"
	PushNotification  NotificationType = "push"
	SMSNotification   NotificationType = "sms"
	SlackNotification NotificationType = "slack"
)

var validNotificationTypes = map[NotificationType]bool{
	EmailNotification: true,
	PushNotification:  true,
	SMSNotification:   true,
	SlackNotification: true,
}

func IsValidNotificationType(s string) bool {
	_, ok := validNotificationTypes[NotificationType(s)]
	return ok
}

type User struct {
	ID                int                `json:"id"`
	Email             string             `json:"email"`
	NotificationTypes []NotificationType `json:"notification_types"`
}
