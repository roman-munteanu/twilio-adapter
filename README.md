twilio-adapter
-----

## Build and run

Run the app:
```
go mod tidy

go mod vendor

go run .\main.go
```

## Resources

* Go quickstart example:
https://www.twilio.com/docs/sms/quickstart/go

* Webhooks:
https://www.twilio.com/docs/usage/webhooks

SMS Webhooks (to check if the message has been delivered successfully):
https://www.twilio.com/docs/usage/webhooks/sms-webhooks

user-defined HTTP callback (webhook) example:
https://www.twilio.com/docs/sms/tutorials/how-to-confirm-delivery-java


* Phone number validation:
https://www.twilio.com/docs/lookup/api

Validate a national phone number
https://www.twilio.com/docs/lookup/tutorials/validation-and-formatting#

* E.164 format
https://support.twilio.com/hc/en-us/articles/223183008-Formatting-International-Phone-Numbers

* The list of country calling codes:
https://en.wikipedia.org/wiki/List_of_country_calling_codes#Alphabetical_listing_by_country_or_region

* Using a list of allowed country codes:
https://www.twilio.com/blog/allow-list-country-code-lookup

* Best practices for phone number validation during new user enrollment
https://www.twilio.com/blog/best-practices-phone-number-validation-user-enrollment

* Trial limitations
https://www.twilio.com/docs/usage/tutorials/how-to-use-your-free-trial-account#trial-account-restrictions-and-limitations

* Scaling with Messaging Services 

https://www.twilio.com/docs/messaging/services

https://www.twilio.com/docs/messaging/guides/best-practices-at-scale

