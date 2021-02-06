package window

type Listener interface {
	Push([]byte)
	Commit() error
}
