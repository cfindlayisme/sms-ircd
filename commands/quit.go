package commands

import (
	"log"
	"strings"

	"github.com/cfindlayisme/sms-ircd/model"
)

// HandleQuit processes the QUIT command
func HandleQuit(client *model.ConnectedUser, command string) {
	// Extract quit message if provided
	parts := strings.SplitN(command, ":", 2)
	quitMsg := "Client quit"
	if len(parts) > 1 {
		quitMsg = parts[1]
	}

	log.Printf("Client %s quit: %s", client.Nick, quitMsg)

	// Send ERROR and close connection
	client.Connection.Write([]byte("ERROR :Closing connection: " + quitMsg + "\r\n"))
	client.Connection.Close()
}
