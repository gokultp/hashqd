package queue

const (
	TubeDefault = "default"
)

var (
	qMap map[string]*Queue
)

func Init() {
	qMap = map[string]*Queue{
		TubeDefault: NewQueue(),
	}
}

func Enqueue(tube string, data []byte) (int, error) {
	if _, found := qMap[tube]; !found {
		qMap[tube] = NewQueue()
	}
	return qMap[tube].Enqueue(data)
}

func Dequeue(tube string, data chan []byte) {
	if _, found := qMap[tube]; !found {
		qMap[tube] = NewQueue()
	}
	qMap[tube].Dequeue(data)
}

func Update(tube string, id int, data []byte) error {
	if _, found := qMap[tube]; !found {
		qMap[tube] = NewQueue()
	}
	return qMap[tube].Update(id, data)
}

func Disconnect(tube string) {
	// if no one is reserved the tube and if is empty
	if q, found := qMap[tube]; found && q.front < q.back {
		delete(qMap, tube)
	}
}
