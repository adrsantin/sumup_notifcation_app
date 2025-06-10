package payments

import (
	"encoding/json"
	"fmt"
	"sumup/notifications/internal/business"
	"sumup/notifications/internal/entities"
	"sumup/notifications/internal/queue"
	"sumup/notifications/internal/repositories"
)

type paymentServiceImpl struct {
	userRepository repositories.UserRepository
	producer       queue.Producer
}

func NewPaymentService(
	userRepository repositories.UserRepository,
	producer queue.Producer,
) business.PaymentService {
	return &paymentServiceImpl{
		userRepository: userRepository,
		producer:       producer,
	}
}

func (s *paymentServiceImpl) ProcessPaymentNotification(userID int, amount float64) error {
	user, err := s.userRepository.GetUserDataByID(userID)
	if err != nil {
		return err
	}
	notificationTypes, err := s.userRepository.GetUserNotificationTypesByUserID(userID)
	if err != nil {
		return err
	}

	for _, notificationType := range notificationTypes {
		message := entities.MessageDTO{
			UserID:           user.ID,
			Name:             user.Name,
			Email:            user.Email,
			Phone:            user.Phone,
			Amount:           amount,
			NotificationType: notificationType,
		}
		messageBytes, err := json.Marshal(message)
		if err != nil {
			return fmt.Errorf("failed to marshal message: %w", err)
		}
		if err := s.producer.Produce(messageBytes); err != nil {
			return fmt.Errorf("failed to produce message: %w", err)
		}
	}
	return nil
}
