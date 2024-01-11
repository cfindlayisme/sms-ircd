package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/cfindlayisme/sms-ircd/ircserver"
	"github.com/cfindlayisme/sms-ircd/model"
)

func main() {
	listener, err := net.Listen("tcp", ":6555")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("IRC server started on port 6555")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	client := model.ConnectedUser{
		IP: net.IPAddr{
			IP: net.ParseIP("127.0.0.1"),
		},
	}
	client.Connection = conn

	ircserver.SendScheduledPings(&client, 5*time.Minute)

	// Read and process client commands
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		// Process client command
		command := string(buf[:n])
		if strings.HasPrefix(strings.ToUpper(command), "NICK") {
			nick := strings.TrimSpace(strings.TrimPrefix(command, "NICK"))
			ircserver.RecieveNick(&client, nick)

			log.Println("Received NICK command:", nick)

		} else if strings.HasPrefix(strings.ToUpper(command), "USER") {
			user := strings.TrimSpace(strings.TrimPrefix(command, "USER"))

			userName := strings.Split(user, " ")
			client.User = userName[0]

			realName := strings.Split(command, ":")
			client.RealName = realName[1]

			log.Println("Received USER command:", user)
			ircserver.SendServerRawMessage(&client, "001", "Welcome to the server "+client.Nick+"!")
			ircserver.SendForceJoin(&client, "#control")

		} else if strings.HasPrefix(strings.ToUpper(command), "JOIN") {
			ircserver.SendRawMessage(&client, "481 JOIN :Clients are not allowed to join channels of their own free will on this IRC server.")

		} else if strings.HasPrefix(strings.ToUpper(command), "PART") {
			ircserver.SendRawMessage(&client, "481 PART :Clients are not allowed to part channels of their own free will on this IRC server.")

		} else if strings.HasPrefix(strings.ToUpper(command), "TOPIC") {
			split2 := strings.Split(command, " ")
			channelName := split2[1]

			if len(split2) == 2 {
				ircserver.SendChannelTopic(&client, channelName)

			} else {
				ircserver.SendRawMessage(&client, "481 TOPIC :Clients are not allowed to adjust topics on this IRC server.")
			}

		} else if strings.HasPrefix(strings.ToUpper(command), "TWILIO") {
			split2 := strings.Split(command, " ")

			if len(split2) != 3 {
				ircserver.SendRawMessage(&client, "461 TWILIO :Invalid TWILIO command. Try /TWILIO <username> <password>")
			} else {
				twilioUsername := split2[1]
				twilioPassword := split2[2]

				client.TwilioUsername = twilioUsername
				client.TwilioPassword = twilioPassword
			}

		} else {
			// TODO: Implement logic to handle other commands

			conn.Write([]byte("ERROR :Unknown command\r\n"))
		}
	}
}
