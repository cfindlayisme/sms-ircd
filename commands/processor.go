package commands

import (
	"strings"

	"github.com/cfindlayisme/sms-ircd/model"
)

// ProcessCommand routes IRC commands to their appropriate handlers
func ProcessCommand(client *model.ConnectedUser, command string) {
	upperCommand := strings.ToUpper(command)

	switch {
	case strings.HasPrefix(upperCommand, "NICK"):
		HandleNick(client, command)
	case strings.HasPrefix(upperCommand, "USER"):
		HandleUser(client, command)
	case strings.HasPrefix(upperCommand, "JOIN"):
		HandleJoin(client, command)
	case strings.HasPrefix(upperCommand, "PART"):
		HandlePart(client, command)
	case strings.HasPrefix(upperCommand, "TOPIC"):
		HandleTopic(client, command)
	case strings.HasPrefix(upperCommand, "TWILIO"):
		HandleTwilio(client, command)
	default:
		HandleUnknown(client, command)
	}
}
