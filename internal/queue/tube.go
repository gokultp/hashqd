package queue

type Tube struct {
	name     string
	queue    *Queue
	reserved int
}
