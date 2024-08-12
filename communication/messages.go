package communication

type Message struct {
	AccountSid          string `json:"account_sid"`
	APIVersion          string `json:"api_version"`
	Body                string `json:"body"`
	DateCreated         string `json:"date_created"`
	DateSent            string `json:"date_sent"`
	Direction           string `json:"direction"`
	ErrorCode           string `json:"error_code"`
	ErrorMessage        string `json:"error_message"`
	From                string `json:"from"`
	NumMedia            string `json:"num_media"`
	NumSegments         string `json:"num_segments"`
	Price               string `json:"price"`
	PriceUnit           string `json:"price_unit"`
	MessagingServiceSid string `json:"messaging_service_sid"`
	SID                 string `json:"sid"`
	Status              string `json:"status"`
	SubresourceUris     string `json:"subresource_uris"`
	Tags                string `json:"tags"`
	To                  string `json:"to"`
	URI                 string `json:"uri"`
}

func NewMessage() *Message {
	return &Message{}
}

// Esse é um modelo de como volta a mensagem da API do Whatsapp
// {
//   "account_sid": "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
//   "api_version": "2010-04-01",
//   "body": "Hello there!",
//   "date_created": "Thu, 24 Aug 2023 05:01:45 +0000",
//   "date_sent": "Thu, 24 Aug 2023 05:01:45 +0000",
//   "date_updated": "Thu, 24 Aug 2023 05:01:45 +0000",
//   "direction": "outbound-api",
//   "error_code": null,
//   "error_message": null,
//   "from": "whatsapp:+14155238886",
//   "num_media": "0",
//   "num_segments": "1",
//   "price": null,
//   "price_unit": null,
//   "messaging_service_sid": "MGaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
//   "sid": "SMaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
//   "status": "queued",
//   "subresource_uris": {
//     "media": "/2010-04-01/Accounts/ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/Messages/SMaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/Media.json"
//   },
//   "tags": {
//     "campaign_name": "Spring Sale 2022",
//     "message_type": "cart_abandoned"
//   },
//   "to": "whatsapp:+15005550006",
//   "uri": "/2010-04-01/Accounts/ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/Messages/SMaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.json"
// }
