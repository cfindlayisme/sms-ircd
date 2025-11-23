package commands

import (
	"log"
	"strings"

	"github.com/cfindlayisme/sms-ircd/ircserver"
	"github.com/cfindlayisme/sms-ircd/model"
)

// HandlePrivmsg processes the PRIVMSG command
func HandlePrivmsg(client *model.ConnectedUser, command string) {
	// PRIVMSG format: PRIVMSG <target> :<message>
	parts := strings.SplitN(command, " ", 3)

	if len(parts) < 3 {
		ircserver.SendRawMessage(client, "461 PRIVMSG :Not enough parameters")
		return
	}

	target := parts[1]
	message := strings.TrimPrefix(parts[2], ":")

	// Check if target is a phone number (starts with +)
	if strings.HasPrefix(target, "+") {
		// TODO: Handle Twilio SMS sending
		log.Printf("PRIVMSG to phone number %s: %s", target, message)
		// For now, just acknowledge receipt
		ircserver.SendRawMessage(client, "NOTICE "+client.Nick+" :Message queued for "+target)
	} else if strings.HasPrefix(target, "#") {
		// Channel message - not allowed for now
		ircserver.SendRawMessage(client, "404 "+target+" :Cannot send to channel")
	} else {
		// Private message to another user - not implemented
		ircserver.SendRawMessage(client, "401 "+target+" :No such nick/channel")
	}
}

// HandleNotice processes the NOTICE command
func HandleNotice(client *model.ConnectedUser, command string) {
	// NOTICE is similar to PRIVMSG but shouldn't trigger automatic replies
	// For now, treat it similarly
	parts := strings.SplitN(command, " ", 3)

	if len(parts) < 3 {
		ircserver.SendRawMessage(client, "461 NOTICE :Not enough parameters")
		return
	}

	target := parts[1]

	if strings.HasPrefix(target, "+") {
		// Phone number target - stub for now
		log.Printf("NOTICE to phone number %s", target)
	}
}
