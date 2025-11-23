package commands

import (
	"github.com/cfindlayisme/sms-ircd/model"
)

// HandleUnknown processes unknown commands
func HandleUnknown(client *model.ConnectedUser, command string) {
	// TODO: Implement logic to handle other commands
	client.Connection.Write([]byte("ERROR :Unknown command\r\n"))
}
