package client

import (
	"log"
	"net"
	"time"

	"github.com/cfindlayisme/sms-ircd/commands"
	"github.com/cfindlayisme/sms-ircd/ircserver"
	"github.com/cfindlayisme/sms-ircd/model"
)

// HandleConnection manages a single client connection
func HandleConnection(conn net.Conn) {
	defer conn.Close()

	client := model.ConnectedUser{
		IP: net.IPAddr{
			IP: net.ParseIP("127.0.0.1"),
		},
	}
	client.Connection = conn

	ircserver.SendScheduledPings(&client, 5*time.Minute)

	// Read and process client commands
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		// Process client command
		command := string(buf[:n])
		commands.ProcessCommand(&client, command)
	}
}
