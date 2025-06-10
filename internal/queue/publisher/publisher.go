package publisher

import (
	"fmt"
	"sumup/notifications/internal/queue"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	KafkaTopic = "notification_topic"
)

type producerImpl struct {
	kafkaProducer *kafka.Producer
}

func NewProducer(
	kafkaProducer *kafka.Producer,
) queue.Producer {
	return &producerImpl{
		kafkaProducer: kafkaProducer,
	}
}

func (p *producerImpl) Produce(message []byte) error {
	topic := KafkaTopic
	err := p.kafkaProducer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: message,
		}, nil,
	)
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}
	return nil
}
