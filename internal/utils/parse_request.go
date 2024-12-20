package utils

import (
	"gateway/internal/models"
	"net/http"
)

func ParseTwilioRequest(r *http.Request) (*models.TwilioMessage, error) {

	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	formValues := r.Form

	data := &models.TwilioMessage{
		MessageSid:          formValues.Get("MessageSid"),
		SmsSid:              formValues.Get("SmsSid"),
		SmsMessageSid:       formValues.Get("SmsMessageSid"),
		AccountSid:          formValues.Get("AccountSid"),
		MessagingServiceSid: formValues.Get("MessagingServiceSid"),
		From:                formValues.Get("From"),
		To:                  formValues.Get("To"),
		Body:                formValues.Get("Body"),
		NumMedia:            formValues.Get("NumMedia"),
		NumSegments:         formValues.Get("NumSegments"),
		ProfileName:         formValues.Get("ProfileName"),
		Wald:                formValues.Get("Wald"),
		Forwarded:           formValues.Get("Forwarded"),
		FrequentlyForwarded: formValues.Get("FrequentlyForwarded"),
		ButtonText:          formValues.Get("ButtonText"),
		Latitute:            formValues.Get("Latitute"),
		Longitude:           formValues.Get("Longitude"),
		Address:             formValues.Get("Address"),
		Label:               formValues.Get("Label"),
	}

	return data, nil
}
