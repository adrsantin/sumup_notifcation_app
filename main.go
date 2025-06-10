package main

import (
	"database/sql"
	"log"
	"net/http"
	"sumup/notifications/internal/api/health"
	paymentsapi "sumup/notifications/internal/api/payments"
	"sumup/notifications/internal/business/notifications"
	paymentsservice "sumup/notifications/internal/business/payments"
	"sumup/notifications/internal/queue"
	"sumup/notifications/internal/queue/publisher"
	"sumup/notifications/internal/queue/worker"
	"sumup/notifications/internal/repositories/users"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
)

const (
	KafkaTopic = "notification_topic"
)

func main() {
	router, worker := load()

	go worker.ProcessMessages()

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}

func load() (*chi.Mux, queue.Worker) {
	router := chi.NewRouter()

	db, err := sql.Open("mysql", "root:password@tcp(go-mysql:3306)/sumup")
	if err != nil {
		log.Fatal(err)
	}

	userRepository := users.NewUserRepository(db)

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "go-kafka:9092",
	})
	if err != nil {
		log.Fatal(err)
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "go-kafka:9092",
		"group.id":          "notification_group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	topic := KafkaTopic
	c.SubscribeTopics([]string{topic}, nil)

	producer := publisher.NewProducer(p)

	paymentService := paymentsservice.NewPaymentService(
		userRepository,
		producer,
	)

	notificationService := notifications.NewNotificationService()

	worker := worker.NewWorker(c, notificationService)

	health.NewHealthAPI(router)
	paymentsapi.NewPaymentsAPI(router, paymentService)

	return router, worker
}
