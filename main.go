package main

import (
	"context"
	"fmt"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioAdapterApp struct {
	ctx    context.Context
	client *twilio.RestClient
}

/*
Set environment variables:
TWILIO_ACCOUNT_SID
TWILIO_AUTH_TOKEN
*/
func main() {
	app := TwilioAdapterApp{
		ctx:    context.Background(),
		client: twilio.NewRestClient(),
	}

	phoneNumberFrom := "+..." // phone number bought on Twilio (Buy a number)
	phoneNumberTo := "+..."   // Vetified Caller ID (my local phone number)

	params := &api.CreateMessageParams{}
	params.SetBody("Test SMS from Roman")
	params.SetFrom(phoneNumberFrom)
	params.SetTo(phoneNumberTo)

	resp, err := app.client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", resp)

	if resp.Sid != nil {
		fmt.Println(*resp.Sid)
	} else {
		fmt.Println(resp.Sid)
	}
}
