package commands

import (
	"log"
	"strings"

	"github.com/cfindlayisme/sms-ircd/ircserver"
	"github.com/cfindlayisme/sms-ircd/model"
)

// HandleNick processes the NICK command
func HandleNick(client *model.ConnectedUser, command string) {
	nick := strings.TrimSpace(strings.TrimPrefix(command, "NICK"))
	ircserver.RecieveNick(client, nick)
	log.Println("Received NICK command:", nick)
}
