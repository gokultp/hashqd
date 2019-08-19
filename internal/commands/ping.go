package commands

import (
	"io"

	"github.com/gokultp/hashqd/internal/session"
)

type Ping struct {
	reader  io.Reader
	session *session.Session
}

func NewPing(s *session.Session, r io.Reader) *Ping {
	return &Ping{
		reader:  r,
		session: s,
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
