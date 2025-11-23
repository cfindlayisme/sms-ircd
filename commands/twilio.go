package commands

import (
	"strings"

	"github.com/cfindlayisme/sms-ircd/ircserver"
	"github.com/cfindlayisme/sms-ircd/model"
)

// HandleTwilio processes the TWILIO command
func HandleTwilio(client *model.ConnectedUser, command string) {
	split2 := strings.Split(command, " ")

	if len(split2) != 3 {
		ircserver.SendRawMessage(client, "461 TWILIO :Invalid TWILIO command. Try /TWILIO <username> <password>")
	} else {
		twilioUsername := split2[1]
		twilioPassword := split2[2]

		client.TwilioUsername = twilioUsername
		client.TwilioPassword = twilioPassword
	}
}
