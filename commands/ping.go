package commands

import (
	"strings"

	"github.com/cfindlayisme/sms-ircd/ircserver"
	"github.com/cfindlayisme/sms-ircd/model"
)

// HandlePing processes the PING command
func HandlePing(client *model.ConnectedUser, command string) {
	// Extract the ping argument
	parts := strings.Fields(command)
	var arg string
	if len(parts) > 1 {
		arg = parts[1]
	}

	// Send PONG response
	if arg != "" {
		ircserver.SendRawMessage(client, "PONG :"+arg)
	} else {
		ircserver.SendRawMessage(client, "PONG")
	}
}

// HandlePong processes the PONG command
func HandlePong(client *model.ConnectedUser, command string) {
	// Client responded to our ping, nothing to do
	// Could update last activity timestamp here if needed
}
