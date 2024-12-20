package models

type BotkitWrapper struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type BotkitMessage struct {
	Type    string `json:"type"`
	Section string `json:"section"`
	Body    string `json:"body"`
}
