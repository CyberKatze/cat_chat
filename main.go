package main

import (
	"flag"

	"github.com/m3dsh/chat/client"
	"github.com/m3dsh/chat/server"
)

var IsServer bool
var Port int

func main() {
	flag.BoolVar(&IsServer, "l", false, "Is server and listen")
	flag.IntVar(&Port, "p", 8888, "Is server and listen")
	flag.Parse()
	if IsServer {
		server.Set(Port)
		server.Start()
	} else {
		client.Set(Port)
		client.Start()
	}
}
