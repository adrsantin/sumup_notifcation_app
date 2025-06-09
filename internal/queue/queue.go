package queue

type Producer interface {
	Produce(message []byte) error
}
