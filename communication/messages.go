package communication

// talvez usar o pacote gorilla/schema seja util ele possui ferramentas para parse de url.values
// esse é o formato que recebemos das mensagens do Twilio
type Message struct {
	SmsMessageSid    string `json:"sms_message_sid"`
	NumMedia         string `json:"num_media"`
	ProfileName      string `json:"profile_name"`
	MessageType      string `json:"message_type"`
	SmsSid           string `json:"sms_sid"`
	WaId             string `json:"wa_id"`
	SmsStatus        string `json:"sms_status"`
	Body             string `json:"body"`
	To               string `json:"to"`
	NumSegments      string `json:"num_segments"`
	ReferralNumMedia string `json:"referral_num_media"`
	MessageSid       string `json:"message_sid"`
	AccountSid       string `json:"account_sid"`
	From             string `json:"from"`
	APIVersion       string `json:"api_version"`
}

// Este é o formato de uma mensagem que vem da Twilio quando recebemos uma mensagem no Bot
// SmsMessageSid    = XXXXXXXXXXXXXXXXXXXXXXXX
// NumMedia         = 0
// ProfileName      = Paulo Campos
// MessageType      = text
// SmsSid           = XXXXXXXXXXXXXXXXXXXXXXXX
// WaId             = 5511944924454
// SmsStatus        = received
// Body             = oi
// To               = whatsapp%!A(MISSING)%!B(MISSING)14155238886
// NumSegments      = 1
// ReferralNumMedia = 0
// MessageSid       = XXXXXXXXXXXXXXXXXXXXXXXX
// AccountSid       = XXXXXXXXXXXXXXXXXXXXXXXX
// From             = whatsapp%!A(MISSING)%!B(MISSING)5511944924454
// ApiVersion       = 2010-04-01
