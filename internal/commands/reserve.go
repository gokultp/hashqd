package commands

import (
	"io"

	"github.com/gokultp/hashqd/internal/queue"
)

type Reserve struct {
	reader io.Reader
}

func NewReserve(reader io.Reader) *Reserve {
	return &Reserve{
		reader: reader,
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
