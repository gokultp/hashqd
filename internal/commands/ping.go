package commands

import (
	"io"
)

type Ping struct {
	reader io.Reader
}

func NewPing(reader io.Reader) *Ping {
	return &Ping{
		reader: reader,
	}
}

func (c *Ping) Decode() *Error {
	return nil
}

func (c *Ping) Exec() (*Response, *Error) {
	return NewResponse("pong\n"), nil
}

func (c *Ping) Command() string {
	return CommandPing
}
