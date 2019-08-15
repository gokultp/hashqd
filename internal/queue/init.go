package queue

import "fmt"

var (
	q *Queue
)

func Init() {
	q = NewQueue()
}

func Enqueue(data []byte) (int, error) {
	return q.Enqueue(data)
}

func Dequeue(data chan []byte) {
	q.Dequeue(data)
}

func Update(id int, data []byte) error {
	return q.Update(id, data)
}

func Log() {
	fmt.Println(q.data, q.front, q.back)
}
