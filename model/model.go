package model

import (
	"net"
)

type ConnectedUser struct {
	Nick           string
	User           string
	RealName       string
	Connection     net.Conn
	Channels       map[string]Channel
	IP             net.IPAddr
	TwilioUsername string
	TwilioPassword string
}

type Channel struct {
	Name  string
	Topic string
	Mode  string
	Users []string
}
