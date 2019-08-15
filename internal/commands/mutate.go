package commands

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/gokultp/hashqd/internal/queue"
)

type Mutate struct {
	data   []byte
	id     int
	reader io.Reader
}

func NewMutate(reader io.Reader) *Mutate {
	return &Mutate{
		reader: reader,
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
			return errUnknown(" on reading request")
		}
		bufSize, err = strconv.Atoi(meta[1])
		if err != nil {
			log.Printf("error on atoi in Mutate.decode %v", err)
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
		log.Printf("Mutate.decode scanner error %v", scanner.Err())
		return errUnknown(" on reading request")

	}
	return errUnknown(" on reading request")
}

func (c *Mutate) Exec() (*Response, *Error) {
	err := queue.Update(c.id, c.data)
	if err != nil {
		return nil, errMutate(err.Error())
	}
	return NewResponse(fmt.Sprintf("%d\n", c.id)), nil
}

func (c *Mutate) Command() string {
	return CommandMutate
}
