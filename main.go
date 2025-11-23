package main

import (
	"log"
	"net"

	"github.com/cfindlayisme/sms-ircd/client"
	"github.com/cfindlayisme/sms-ircd/env"
)

func main() {
	listener, err := net.Listen("tcp", ":"+env.GetServerPort())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("IRC server started on port", env.GetServerPort())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go client.HandleConnection(conn)
	}
}
