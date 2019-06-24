package queue

type IQueue interface {
	Enqueue([]byte) (int, error)
	Dequeue() ([]byte, error)
	Update(int, []byte) error
}
