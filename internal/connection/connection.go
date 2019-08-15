package connection

import (
	"bufio"
	"log"
	"net"

	"github.com/gokultp/hashqd/internal/commands"
)

func Listen(port string) error {
	l, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		return err
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}
		go hanldeConn(c)
	}
}

func hanldeConn(c net.Conn) {
	for {
		reader := bufio.NewReader(c)
		cmd, err := commands.GetCommand(reader)
		if err != nil {
			log.Print(err)
			return
		}
		if cerr := cmd.Decode(); cerr != nil {
			c.Write(cerr.Bytes())
			continue
		}
		res, cerr := cmd.Exec()
		if cerr != nil {
			c.Write(cerr.Bytes())
			continue
		}
		c.Write(res.Bytes())
		if cmd.Command() == commands.CommandFin {
			break
		}
	}
	c.Close()
}
