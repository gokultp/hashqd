package commands

import (
	"io"
)

type Fin struct {
	reader io.Reader
}

func NewFin(reader io.Reader) *Fin {
	return &Fin{
		reader: reader,
	}
}

func (c *Fin) Decode() *Error {
	return nil
}

func (c *Fin) Exec() (*Response, *Error) {
	return NewResponse("bye\n"), nil
}

func (c *Fin) Command() string {
	return CommandFin
}
