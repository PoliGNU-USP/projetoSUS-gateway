package models

import "encoding/xml"

// Baseado em
// https://www.twilio.com/docs/messaging/guides/webhook-request#parameters-in-twilios-request-to-your-application

type TwilioMessage struct {
	MessageSid          string `form:"MessageSid"`
	SmsSid              string `form:"SmsSid"`
	SmsMessageSid       string `form:"SmsMessageSid"`
	AccountSid          string `form:"AccountSid"`
	MessagingServiceSid string `form:"MessagingServiceSid"`
	From                string `form:"From"`
	To                  string `form:"To"`
	Body                string `form:"Body"`
	NumMedia            string `form:"NumMedia"`
	NumSegments         string `form:"NumSegments"`

	ProfileName         string `form:"ProfileName"`
	Wald                string `form:"Wald"`
	Forwarded           string `form:"Forwarded"`
	FrequentlyForwarded string `form:"FrequentlyForwarded"`
	ButtonText          string `form:"ButtonText"`
	Latitute            string `form:"Latitute"`
	Longitude           string `form:"Longitude"`
	Address             string `form:"Address"`
	Label               string `form:"Label"`
}

// Essa parte é bem confusa por causa desse jeito diferente da Twilio de responder as mensagens
// A documentação deles é bem ruim também
// https://www.twilio.com/docs/messaging/twiml
type TwiML_Message struct {
	Body string `xml:",chardata"`
}

type TwiML struct {
	XMLName  xml.Name        `xml:"Response"`
	Messages []TwiML_Message `xml:"Message"`
}
