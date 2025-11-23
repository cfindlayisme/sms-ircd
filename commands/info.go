package commands

import (
	"strings"

	"github.com/cfindlayisme/sms-ircd/ircserver"
	"github.com/cfindlayisme/sms-ircd/model"
)

// HandleMode processes the MODE command
func HandleMode(client *model.ConnectedUser, command string) {
	parts := strings.Fields(command)

	if len(parts) < 2 {
		ircserver.SendRawMessage(client, "461 MODE :Not enough parameters")
		return
	}

	target := parts[1]

	if strings.HasPrefix(target, "#") {
		// Channel mode query
		if len(parts) == 2 {
			// Query channel modes
			ircserver.SendRawMessage(client, "324 "+client.Nick+" "+target+" +nt")
			ircserver.SendRawMessage(client, "329 "+client.Nick+" "+target+" 0")
		} else {
			// Setting channel modes - not allowed
			ircserver.SendRawMessage(client, "482 "+target+" :You're not channel operator")
		}
	} else if target == client.Nick {
		// User mode query/set
		if len(parts) == 2 {
			// Query user modes
			ircserver.SendRawMessage(client, "221 "+client.Nick+" +i")
		} else {
			// Setting user modes - accept but don't actually change anything
			ircserver.SendRawMessage(client, "MODE "+client.Nick+" "+parts[2])
		}
	} else {
		ircserver.SendRawMessage(client, "502 :Can't change mode for other users")
	}
}

// HandleWho processes the WHO command
func HandleWho(client *model.ConnectedUser, command string) {
	parts := strings.Fields(command)

	if len(parts) < 2 {
		ircserver.SendRawMessage(client, "461 WHO :Not enough parameters")
		return
	}

	target := parts[1]

	if strings.HasPrefix(target, "#") {
		// WHO for channel - return basic info
		// Format: 352 <client> <channel> <user> <host> <server> <nick> <H|G> :<hopcount> <realname>
		ircserver.SendRawMessage(client, "352 "+client.Nick+" "+target+" "+client.User+" localhost server "+client.Nick+" H :0 "+client.RealName)
		ircserver.SendRawMessage(client, "315 "+client.Nick+" "+target+" :End of WHO list")
	} else {
		// WHO for user
		ircserver.SendRawMessage(client, "315 "+client.Nick+" "+target+" :End of WHO list")
	}
}

// HandleWhois processes the WHOIS command
func HandleWhois(client *model.ConnectedUser, command string) {
	parts := strings.Fields(command)

	if len(parts) < 2 {
		ircserver.SendRawMessage(client, "461 WHOIS :Not enough parameters")
		return
	}

	target := parts[1]

	if target == client.Nick {
		// WHOIS for self
		ircserver.SendRawMessage(client, "311 "+client.Nick+" "+target+" "+client.User+" localhost * :"+client.RealName)
		ircserver.SendRawMessage(client, "312 "+client.Nick+" "+target+" server :SMS-IRC Gateway")
		ircserver.SendRawMessage(client, "318 "+client.Nick+" "+target+" :End of WHOIS list")
	} else {
		// No such user
		ircserver.SendRawMessage(client, "401 "+target+" :No such nick/channel")
		ircserver.SendRawMessage(client, "318 "+client.Nick+" "+target+" :End of WHOIS list")
	}
}

// HandleList processes the LIST command
func HandleList(client *model.ConnectedUser, command string) {
	// Return list of channels
	ircserver.SendRawMessage(client, "321 "+client.Nick+" Channel :Users Name")
	ircserver.SendRawMessage(client, "322 "+client.Nick+" #control 1 :Control channel for SMS gateway")
	ircserver.SendRawMessage(client, "323 "+client.Nick+" :End of LIST")
}

// HandleNames processes the NAMES command
func HandleNames(client *model.ConnectedUser, command string) {
	parts := strings.Fields(command)

	var channel string
	if len(parts) > 1 {
		channel = parts[1]
	} else {
		channel = "#control"
	}

	// Return names list for channel
	ircserver.SendRawMessage(client, "353 "+client.Nick+" = "+channel+" :"+client.Nick)
	ircserver.SendRawMessage(client, "366 "+client.Nick+" "+channel+" :End of NAMES list")
}
