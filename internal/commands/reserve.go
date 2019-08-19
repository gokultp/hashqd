package commands

import (
	"io"

	"github.com/gokultp/hashqd/internal/queue"
	"github.com/gokultp/hashqd/internal/session"
)

type Reserve struct {
	reader  io.Reader
	session *session.Session
}

func NewReserve(s *session.Session, r io.Reader) *Reserve {
	return &Reserve{
		reader:  r,
		session: s,
	}
}

func (c *Reserve) Decode() *Error {
	return nil
}

func (c *Reserve) Exec() (*Response, *Error) {
	dataChan := make(chan []byte)
	go queue.Dequeue(dataChan)
	res := NewResponse()
	for {
		select {
		case data := <-dataChan:
			res.SetBytes(data)
			return res, nil
		}
	}

}

func (c *Reserve) Command() string {
	return CommandReserve
}
