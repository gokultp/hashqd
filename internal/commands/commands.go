package commands

import (
	"errors"
	"io"
)

const (
	CommandPut     = "put"
	CommandPing    = "ping"
	CommandFin     = "fin"
	CommandReserve = "reserve"
	CommandMutate  = "mutate"
)

func GetCommand(r io.Reader) (ICommand, error) {
	commandName, err := readCommand(r)
	if err != nil {
		return nil, err
	}
	switch commandName {
	case CommandPut:
		return NewPut(r), nil
	case CommandPing:
		return NewPing(r), nil
	case CommandFin:
		return NewFin(r), nil
	case CommandReserve:
		return NewReserve(r), nil
	case CommandMutate:
		return NewMutate(r), nil
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
