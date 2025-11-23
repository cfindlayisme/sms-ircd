package commands

import (
	"log"
	"strings"

	"github.com/cfindlayisme/sms-ircd/ircserver"
	"github.com/cfindlayisme/sms-ircd/model"
)

// HandleUser processes the USER command
func HandleUser(client *model.ConnectedUser, command string) {
	user := strings.TrimSpace(strings.TrimPrefix(command, "USER"))

	userName := strings.Split(user, " ")
	client.User = userName[0]

	realName := strings.Split(command, ":")
	client.RealName = realName[1]

	log.Println("Received USER command:", user)
	ircserver.SendServerRawMessage(client, "001", "Welcome to the server "+client.Nick+"!")
	ircserver.SendForceJoin(client, "#control")
}
