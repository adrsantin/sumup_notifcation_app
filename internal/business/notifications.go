package business

import (
	"fmt"
	"sumup/notifications/internal/repositories"
)

type notificationServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewNotificationService(
	userRepository repositories.UserRepository,
) NotificationService {
	return &notificationServiceImpl{
		userRepository: userRepository,
	}
}

func (s *notificationServiceImpl) SendPaymentNotifications(userID int, amount float64) error {
	user, err := s.userRepository.GetUserDataByID(userID)
	if err != nil {
		return err
	}
	notificationTypes, err := s.userRepository.GetUserNotificationTypesByUserID(userID)
	if err != nil {
		return err
	}
	user.NotificationTypes = notificationTypes

	// TODO: implement sending notifications logic
	fmt.Println(user)
	return nil
}
