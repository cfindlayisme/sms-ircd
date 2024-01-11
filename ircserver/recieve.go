package ircserver

import (
	"log"

	"github.com/cfindlayisme/sms-ircd/model"
)

func RecievePrivmsg(client *model.ConnectedUser, target string, message string) {

}

func RecieveNick(client *model.ConnectedUser, newNick string) {
	client.Nick = newNick
}

func RecieveJoin(client *model.ConnectedUser, channelName string) {
	channel := model.Channel{}
	channel.Name = channelName
	channel.Mode = "nt"

	channelToUpdate, ok := client.Channels[channelName]
	if ok {
		log.Println(client.Nick, ": User", client.Nick, "already in channel "+channelName, "that they asked to join again")
		return
	}

	if client.Channels == nil {
		client.Channels = make(map[string]model.Channel)
	}

	// Update the Users field
	channelToUpdate.Users = append(channelToUpdate.Users, client.Nick)

	// Assign the updated channel back to the map
	client.Channels[channelName] = channelToUpdate

	log.Println(client.Nick, ": joined channel "+channelName)
	SendRawMessage(client, ":"+client.Nick+"!"+client.User+"@"+client.IP.String()+" JOIN :"+channelName)
	SendChannelTopic(client, channelName)
	SendChannelUsersList(client, channelName)
}
