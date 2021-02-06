package window

import (
	"time"
)

type Window struct {
	C        chan []byte
	listener Listener
	maxSize  int16
	size     int16
	ticker   *ticker
}

func New(listener Listener, size int16, ttl time.Duration) *Window {
	return &Window{
		C:        make(chan []byte),
		listener: listener,
		maxSize:  size,
		ticker:   NewTicker(ttl),
	}
}

func (w *Window) Exec() {

	for {
		select {

		case data := <-w.C:
			w.listener.Push(data)
			w.size++
			if w.size == w.maxSize {
				w.listener.Commit()
				w.ticker.Reset()
				w.size = 0
			}
		case <-w.ticker.C:
			w.listener.Commit()
			w.ticker.Reset()
			w.size = 0
		}
	}
}
