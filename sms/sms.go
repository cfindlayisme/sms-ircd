package sms

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cfindlayisme/sms-ircd/model"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSms(client *model.ConnectedUser, from string, to string, body string) {
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: client.TwilioUsername,
		Password: client.TwilioPassword,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(body)

	resp, err := twilioClient.Api.CreateMessage(params)
	if err != nil {
		log.Println("Error sending SMS message:", err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Sent SMS Message. Got response:", string(response))
	}
}
