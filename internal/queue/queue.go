package queue

//go:generate mockgen -source=queue.go -destination=../mocks/queue.go -package=mocks
type Producer interface {
	Produce(message []byte) error
}

type Worker interface {
	ProcessMessages()
}
