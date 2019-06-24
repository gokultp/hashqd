package commands

import "io"

type Put struct {
	data   []byte
	reader io.Reader
}

func NewPut(reader io.Reader) *Put {
	return &Put{
		reader: reader,
	}
}

func (c *Put) Decode() error {
	return nil
}

func (c *Put) Exec() error {
	return nil
}

func (c *Put) Command() string {
	return CommandPut
}
