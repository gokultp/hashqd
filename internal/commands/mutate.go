package commands

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/gokultp/hashqd/internal/queue"
	"github.com/gokultp/hashqd/internal/session"
)

type Mutate struct {
	data    []byte
	id      int
	reader  io.Reader
	session *session.Session
}

func NewMutate(s *session.Session, r io.Reader) *Mutate {
	return &Mutate{
		reader:  r,
		session: s,
	}
}

func (c *Mutate) Decode() *Error {
	var bufSize int
	var err error
	scanner := bufio.NewScanner(c.reader)
	scanner.Split(bufio.ScanLines)

	if scanner.Scan() {
		strMeta := scanner.Text()
		meta := strings.Split(strMeta, " ")
		c.id, err = strconv.Atoi(meta[0])
		if err != nil {
			log.Printf("error on atoi in Mutate.decode %v", err)
			return errMutate("id should be a string")
		}
		bufSize, err = strconv.Atoi(meta[1])
		if err != nil {
			log.Printf("error on atoi in Mutate.decode %v", err)
			return errMutate("invalid content length")
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
		log.Printf("Mutate.decode scanner error %v", scanner.Err())
		return errBadRequest("bad content")

	}
	return errUnknown("reading request")
}

func (c *Mutate) Exec() (*Response, *Error) {
	err := queue.Update(c.session.Tube, c.id, c.data)
	if err != nil {
		log.Printf("error on mutation %v", err)
		return nil, errMutate(err.Error())
	}
	return NewResponse(fmt.Sprintf("%d\n", c.id)), nil
}

func (c *Mutate) Command() string {
	return CommandMutate
}
