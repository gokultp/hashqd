package commands

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/gokultp/hashqd/internal/queue"
	"github.com/gokultp/hashqd/internal/session"
)

type Put struct {
	session *session.Session
	data    []byte
	reader  io.Reader
}

func NewPut(s *session.Session, r io.Reader) *Put {
	return &Put{
		session: s,
		reader:  r,
	}
}

func (c *Put) Decode() *Error {
	var bufSize int
	var err error
	scanner := bufio.NewScanner(c.reader)
	scanner.Split(bufio.ScanLines)

	if scanner.Scan() {
		strSize := scanner.Text()
		bufSize, err = strconv.Atoi(strSize)
		if err != nil {
			log.Printf("error on atoi in put.decode %v", err)
			return errUnknown(" on reading request")
		}
	}
	if scanner.Scan() {
		data := scanner.Bytes()
		if len(data) != bufSize {
			return errBadRequest("invalid content length")
		}
		c.data = data
		return nil
	}
	if scanner.Err() != nil {
		log.Printf("put.decode scanner error %v", scanner.Err())
		return errUnknown(" on reading request")

	}
	return errUnknown(" on reading request")
}

func (c *Put) Exec() (*Response, *Error) {
	id, err := queue.Enqueue(c.session.Tube, c.data)
	if err != nil {
		return nil, errPut(err.Error())
	}
	return NewResponse(fmt.Sprintf("%d\n", id)), nil
}

func (c *Put) Command() string {
	return CommandPut
}
