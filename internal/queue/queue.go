package queue

import (
	"fmt"
	"sync"
)

// Queue is a simple in-memory queue implementation
type Queue struct {
	data  map[int][]byte
	front int
	back  int
	lock  sync.Mutex
	dlock sync.Mutex
}

// NewQueue will return a new instance of Queue
func NewQueue() *Queue {
	return &Queue{
		data:  map[int][]byte{},
		front: -1,
		back:  0,
	}
}

// Enqueue inserts data to queue
func (q *Queue) Enqueue(data []byte) (int, error) {
	q.lock.Lock()
	q.data[q.front+1] = data
	f := q.front + 1
	q.front = f
	q.lock.Unlock()
	return q.front, nil

}

// Dequeue dequeues data
func (q *Queue) Dequeue(data chan []byte) {
	q.dlock.Lock()
	for q.front < q.back {
	}
	q.lock.Lock()
	d := q.data[q.back]
	delete(q.data, q.back)
	q.back++
	q.lock.Unlock()
	q.dlock.Unlock()
	data <- d
	return
}

// Update will update an already enqueued job
func (q *Queue) Update(id int, data []byte) error {
	if _, ok := q.data[id]; !ok {
		return fmt.Errorf("could not find any job with id : %d", id)
	}
	q.lock.Lock()
	q.data[id] = data
	q.lock.Unlock()
	return nil
}

// Count will return number of active items in the queue
func (q *Queue) Count() int {
	return q.front - q.back + 1
}
