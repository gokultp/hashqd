package main

import (
	"github.com/gokultp/hashqd/internal/connection"
	"github.com/gokultp/hashqd/internal/queue"
)

func main() {
	queue.Init()
	err := connection.Listen("9000")
	if err != nil {
		panic(err)
	}
}
