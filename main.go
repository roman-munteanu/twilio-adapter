package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioAdapterApp struct {
	ctx                 context.Context
	client              *twilio.RestClient
	verifyServiceSID    string
	messagingServiceSID string
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

	verifyServiceSID := os.Getenv("TWILIO_VERIFY_SERVICE_SID")
	if verifyServiceSID == "" {
		fmt.Println("Twilio Verification Service SID not found")
		return
	}
	app.verifyServiceSID = verifyServiceSID

	messagingServiceSID := os.Getenv("TWILIO_MESSAGING_SERVICE")
	if messagingServiceSID == "" {
		fmt.Println("Twilio Messaging Service SID not found")
		return
	}
	app.messagingServiceSID = messagingServiceSID

	// app.sendMessage()

	// app.verify()

	// app.checkVerificationToken()

	app.sendWithMessagingService()
}

func (app *TwilioAdapterApp) sendMessage() {
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

func (app *TwilioAdapterApp) verify() {
	phoneNumberTo := "+..." // my local phone number

	params := &verify.CreateVerificationParams{}
	params.SetTo(phoneNumberTo)
	params.SetChannel("sms")

	resp, err := app.client.VerifyV2.CreateVerification(app.verifyServiceSID, params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", resp)

	if resp.Status != nil {
		fmt.Println(*resp.Status)
	} else {
		fmt.Println(resp.Status)
	}
}

func (app *TwilioAdapterApp) checkVerificationToken() {
	var token string
	flag.StringVar(&token, "t", "", "enter verification code")
	flag.Parse()

	if token == "" {
		fmt.Println("please specify the verification code")
		return
	}
	fmt.Println("token:", token)

	phoneNumberTo := "+..." // my local phone number

	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phoneNumberTo)
	params.SetCode(token)

	resp, err := app.client.VerifyV2.CreateVerificationCheck(app.verifyServiceSID, params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.Status != nil {
		fmt.Println(*resp.Status)
	} else {
		fmt.Println(resp.Status)
	}
}

func (app *TwilioAdapterApp) sendWithMessagingService() {
	phoneNumberTo := "+..." // my local phone number

	params := &api.CreateMessageParams{}
	params.SetBody("Test SMS from Roman with Messaging Service")
	params.SetMessagingServiceSid(app.messagingServiceSID)
	params.SetTo(phoneNumberTo)

	resp, err := app.client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", resp)

	if resp.Sid != nil {
		fmt.Println(*resp.Sid)
	} else {
		fmt.Println(resp.Sid)
	}
}
