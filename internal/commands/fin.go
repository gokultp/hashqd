package commands

import (
	"io"

	"github.com/gokultp/hashqd/internal/session"
)

type Fin struct {
	reader  io.Reader
	session *session.Session
}

func NewFin(s *session.Session, r io.Reader) *Fin {
	return &Fin{
		reader:  r,
		session: s,
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
