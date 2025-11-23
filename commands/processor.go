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
	case strings.HasPrefix(upperCommand, "PING"):
		HandlePing(client, command)
	case strings.HasPrefix(upperCommand, "PONG"):
		HandlePong(client, command)
	case strings.HasPrefix(upperCommand, "JOIN"):
		HandleJoin(client, command)
	case strings.HasPrefix(upperCommand, "PART"):
		HandlePart(client, command)
	case strings.HasPrefix(upperCommand, "QUIT"):
		HandleQuit(client, command)
	case strings.HasPrefix(upperCommand, "PRIVMSG"):
		HandlePrivmsg(client, command)
	case strings.HasPrefix(upperCommand, "NOTICE"):
		HandleNotice(client, command)
	case strings.HasPrefix(upperCommand, "MODE"):
		HandleMode(client, command)
	case strings.HasPrefix(upperCommand, "WHO"):
		HandleWho(client, command)
	case strings.HasPrefix(upperCommand, "WHOIS"):
		HandleWhois(client, command)
	case strings.HasPrefix(upperCommand, "LIST"):
		HandleList(client, command)
	case strings.HasPrefix(upperCommand, "NAMES"):
		HandleNames(client, command)
	case strings.HasPrefix(upperCommand, "TOPIC"):
		HandleTopic(client, command)
	case strings.HasPrefix(upperCommand, "TWILIO"):
		HandleTwilio(client, command)
	default:
		HandleUnknown(client, command)
	}
}
