package commands

import (
	"bufio"
	"errors"
	"io"
)

const (
	CommandPut = "put"
)

func GetCommand(r io.Reader) (ICommand, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		commandName := scanner.Text()
		switch commandName {
		case CommandPut:
			return NewPut(r), nil
		}
	}
	return nil, errors.New("invalid command, valid commands are <put, ...>")
}
