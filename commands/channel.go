package commands

import (
	"strings"

	"github.com/cfindlayisme/sms-ircd/ircserver"
	"github.com/cfindlayisme/sms-ircd/model"
)

// HandleJoin processes the JOIN command
func HandleJoin(client *model.ConnectedUser, command string) {
	ircserver.SendRawMessage(client, "481 JOIN :Clients are not allowed to join channels of their own free will on this IRC server.")
}

// HandlePart processes the PART command
func HandlePart(client *model.ConnectedUser, command string) {
	ircserver.SendRawMessage(client, "481 PART :Clients are not allowed to part channels of their own free will on this IRC server.")
}

// HandleTopic processes the TOPIC command
func HandleTopic(client *model.ConnectedUser, command string) {
	split2 := strings.Split(command, " ")
	channelName := split2[1]

	if len(split2) == 2 {
		ircserver.SendChannelTopic(client, channelName)
	} else {
		ircserver.SendRawMessage(client, "481 TOPIC :Clients are not allowed to adjust topics on this IRC server.")
	}
}
