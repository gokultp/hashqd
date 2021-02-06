package storage

import (
	"bufio"
	"os"
)

type Listener struct {
	data []byte
	file *os.File
}

func NewListener() *Listener {
	return &Listener{}
}

func (l *Listener) Push(data []byte) {
	l.data = append(l.data, '\n')
	l.data = append(l.data, data...)
}

func (l *Listener) Commit() (err error) {
	if l.file == nil {
		l.file, err = os.OpenFile("/tmp/testq", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
	}
	b := bufio.NewWriterSize(l.file, 128)
	n, err := b.Write(l.data)
	if err != nil {
		return
	}
	if n != len(l.data) {
		return
	}
	l.data = []byte{}
	return
}
