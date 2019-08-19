package commands

import (
	"errors"
	"io"

	"github.com/gokultp/hashqd/internal/session"
)

const (
	CommandPut     = "put"
	CommandPing    = "ping"
	CommandFin     = "fin"
	CommandReserve = "reserve"
	CommandMutate  = "mutate"
	CommandWatch   = "watch"
)

func GetCommand(s *session.Session, r io.Reader) (ICommand, error) {
	commandName, err := readCommand(r)
	if err != nil {
		return nil, err
	}
	switch commandName {
	case CommandPut:
		return NewPut(s, r), nil
	case CommandPing:
		return NewPing(s, r), nil
	case CommandFin:
		return NewFin(s, r), nil
	case CommandReserve:
		return NewReserve(s, r), nil
	case CommandMutate:
		return NewMutate(s, r), nil
	case CommandWatch:
		return NewWatch(s, r), nil
	}
	return nil, errors.New("invalid command, valid commands are <put, ping, fin, reserve...>")
}

func readCommand(r io.Reader) (string, error) {
	command := ""
	buffer := make([]byte, 1)
	for {
		_, err := r.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if buffer[0] == ' ' || buffer[0] == '\n' || buffer[0] == '\r' {
			break
		}
		command = command + string(buffer[0])
	}
	return command, nil
}
