package ircserver

import (
	"log"
	"time"

	"github.com/cfindlayisme/sms-ircd/env"
	"github.com/cfindlayisme/sms-ircd/model"
)

var serverName = env.GetServerName()

func SendServerRawMessage(client *model.ConnectedUser, code string, message string) {
	client.Connection.Write([]byte(":" + serverName + " " + code + " " + client.Nick + " :" + message + "\r\n"))
}

func SendRawMessage(client *model.ConnectedUser, message string) {
	client.Connection.Write([]byte(message + "\r\n"))
}

func SendServerPrivmsg(client *model.ConnectedUser, from string, message string) {
	client.Connection.Write([]byte(":" + from + " PRIVMSG " + client.Nick + " :" + message + "\r\n"))
}

func SendChannelTopic(client *model.ConnectedUser, channelName string) {
	client.Connection.Write([]byte(":" + serverName + " 332 " + client.Nick + " " + channelName + " :" + client.Channels[channelName].Topic + "\r\n"))
}

func SendChannelUsersList(client *model.ConnectedUser, channelName string) {
	users := client.Channels[channelName].Users

	for i := range users {

		client.Connection.Write([]byte(":" + serverName + " 353 " + client.Nick + " = " + channelName + " :" + users[i] + "\r\n"))
	}

	client.Connection.Write([]byte(":" + serverName + " 366 " + client.Nick + " " + channelName + " :End of /NAMES list.\r\n"))
}

func SetTopic(client *model.ConnectedUser, channelName string, topic string) {
	channelToUpdate := client.Channels[channelName]
	channelToUpdate.Topic = topic

	client.Channels[channelName] = channelToUpdate

	SendChannelTopic(client, channelName)
	log.Println(client.Nick, ": Set topic for channel", channelName, "to", topic)
}

func SendToClientPrivmsg(client *model.ConnectedUser, from string, message string) {
	client.Connection.Write([]byte(":" + from + " PRIVMSG " + client.Nick + " :" + message + "\r\n"))
}

func SendScheduledPings(client *model.ConnectedUser, interval time.Duration) {
	ticker := time.NewTicker(interval)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				if client.Connection == nil { // fix this up for net.conn TODO
					close(quit)
					return
				}
				client.Connection.Write([]byte("PING :" + serverName + "\r\n"))
				log.Println(client.Nick, ": Sent PING to", client.Nick)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func SendForceJoin(client *model.ConnectedUser, channelName string) {
	RecieveJoin(client, channelName)
}
