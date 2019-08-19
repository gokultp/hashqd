package main

import (
	"flag"

	"github.com/gokultp/hashqd/internal/connection"
	"github.com/gokultp/hashqd/internal/queue"
	"github.com/gokultp/hashqd/internal/version"
)

func main() {
	versionFound := false
	configPath := "/etc/hashqd/hashqd.conf"
	flag.BoolVar(&versionFound, "version", false, "prints version details")
	flag.StringVar(&configPath, "config-path", "/etc/hashqd/hashqd.conf", "path of the config file")
	flag.Parse()
	if versionFound {
		version.PrintVersion()
		return
	}
	queue.Init()
	err := connection.Listen("9000")
	if err != nil {
		panic(err)
	}
}
