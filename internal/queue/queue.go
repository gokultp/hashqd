package queue

import (
	"errors"
	"fmt"
	"sync"
)

type Queue struct {
	data  map[int][]byte
	front int
	back  int
	lock  sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		data:  map[int][]byte{},
		front: -1,
		back:  0,
	}
}

func (q *Queue) Enqueue(data []byte) (int, error) {
	q.lock.Lock()
	q.data[q.front+1] = data
	f := q.front + 1
	q.front = f
	q.lock.Unlock()
	return q.front, nil

}

func (q *Queue) Dequeue(data chan []byte) {
	for q.front < q.back {
	}
	q.lock.Lock()
	d := q.data[q.back]
	delete(q.data, q.back)
	q.back++
	q.lock.Unlock()
	data <- d
	return
}

func (q *Queue) Update(id int, data []byte) error {
	if _, ok := q.data[id]; !ok {
		return errors.New(fmt.Sprintf("could not find any job with id : %d", id))
	}
	q.lock.Lock()
	q.data[id] = data
	q.lock.Unlock()
	return nil
}
