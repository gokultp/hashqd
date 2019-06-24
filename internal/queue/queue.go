package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	data  map[int][]byte
	front int
	back  int
}

func NewQueue() *Queue {
	return &Queue{
		data:  map[int][]byte{},
		front: -1,
		back:  0,
	}
}

func (q *Queue) Enqueue(data []byte) (int, error) {
	q.data[q.front+1] = data
	q.front += 1
	return q.front, nil
}

func (q *Queue) Dequeue() ([]byte, error) {
	if q.front < q.back {
		return nil, errors.New("Queue is empty")
	}
	data := q.data[q.back]
	delete(q.data, q.back)
	q.back--
	return data, nil
}

func (q *Queue) Update(id int, data []byte) error {
	if _, ok := q.data[id]; !ok {
		return errors.New(fmt.Sprintf("could not find any job with id : %d", id))
	}
	q.data[id] = data
	return nil
}
