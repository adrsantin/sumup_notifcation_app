package notifications

import (
	"fmt"
	"sumup/notifications/internal/business"
	"sumup/notifications/internal/entities"
)

type notificationServiceImpl struct {
}

func NewNotificationService() business.NotificationService {
	return &notificationServiceImpl{}
}

func (s *notificationServiceImpl) SendNotification(user *entities.User, amount float64, notificationType entities.NotificationType) error {
	f, ok := NotificationFuncions[notificationType]
	if ok {
		return f(user, amount)
	}
	return fmt.Errorf("notification type %s not supported", notificationType)
}

type NotificationFunction func(user *entities.User, amount float64) error

var NotificationFuncions = map[entities.NotificationType]NotificationFunction{
	entities.EmailNotification: func(user *entities.User, amount float64) error {
		// Implement email notification logic here
		fmt.Println("Sending email notification to:", user.Email, "for amount:", amount)
		return nil
	},
	entities.SMSNotification: func(user *entities.User, amount float64) error {
		// Implement SMS notification logic here
		fmt.Println("Sending sms notification to:", user.Phone, "for amount:", amount)
		return nil
	},
}
