package main

import (
	"log"
	"os"

	"github.com/go-mangos/mangos/protocol/rep"
	//"github.com/go-mangos/mangos/protocol/req"
	"github.com/go-mangos/mangos/transport/tcp"
)

func main() {

	sock, err := rep.NewSocket()

	if err != nil {
		log.Fatalf("Can't get new response socker: %s\n", err)
		os.Exit(1)
	}

	sock.AddTransport(tcp.NewTransport())

	if err = sock.Listen(":88"); err != nil {
		log.Fatalf("Can't get new response socker: %s\n", err)
	}

}
