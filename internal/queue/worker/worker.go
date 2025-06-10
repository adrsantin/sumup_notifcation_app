package worker

import (
	"encoding/json"
	"fmt"
	"sumup/notifications/internal/business"
	"sumup/notifications/internal/entities"
	"sumup/notifications/internal/queue"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type workerImpl struct {
	kafkaConsumer       *kafka.Consumer
	notificationService business.NotificationService
}

func NewWorker(
	kafkaConsumer *kafka.Consumer,
	notificationService business.NotificationService,
) queue.Worker {
	return &workerImpl{
		kafkaConsumer:       kafkaConsumer,
		notificationService: notificationService,
	}
}

func (c *workerImpl) ProcessMessages() {
	for {
		msg, err := c.kafkaConsumer.ReadMessage(-1)
		if err == nil {
			var message entities.MessageDTO
			err := json.Unmarshal(msg.Value, &message)
			if err != nil {
				fmt.Printf("Error decoding message: %v\n", err)
				continue
			}

			c.notificationService.SendNotification(
				&entities.User{
					ID:    message.UserID,
					Name:  message.Name,
					Email: message.Email,
					Phone: message.Phone,
				},
				message.Amount,
				message.NotificationType,
			)

		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
