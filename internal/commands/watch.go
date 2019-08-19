package commands

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/gokultp/hashqd/internal/session"
)

type Watch struct {
	reader   io.Reader
	session  *session.Session
	tubename string
}

func NewWatch(s *session.Session, r io.Reader) *Watch {
	return &Watch{
		reader:  r,
		session: s,
	}
}

func (c *Watch) Decode() *Error {
	scanner := bufio.NewScanner(c.reader)
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		c.tubename = scanner.Text()
		return nil
	}
	if scanner.Err() != nil {
		log.Printf("put.decode scanner error %v", scanner.Err())
		return errUnknown(" on reading request")

	}
	return errUnknown(" on reading request")
}

func (c *Watch) Exec() (*Response, *Error) {
	c.session.SetTube(c.tubename)
	return NewResponse(fmt.Sprintf("watching %s\n", c.tubename)), nil
}

func (c *Watch) Command() string {
	return CommandWatch
}
