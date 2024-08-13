package communication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type BotkitMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type BotkitInnerMessage struct {
	Type    string `json:"type"`
	Section string `json:"section"`
	Body    string `json:"body"`
}

func SendToBotKit(msg Message) ([]string, error) {

	url := os.Getenv("BOTKIT_WEBHOOK")

	var reply []string

	payload := map[string]string{
		"type":    "message",
		"text":    msg.Body,
		"channel": "webhook",
		"user":    msg.From,
	}

	// Marshal no payload para JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return reply, fmt.Errorf("Error marshaling JSON: %v", err)
	}

	// Criando uma requicao POST com o Json de Payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return reply, fmt.Errorf("Error creating request: %v", err)
	}

	// Definindo o tipo de conteudo para application/json
	req.Header.Set("Content-Type", "application/json")

	// Util para debuggar
	log.Printf("Payload: %s", string(jsonData))
	for key, value := range req.Header {
		log.Printf("Header: %s: %s", key, value)
	}

	// Enviando a requisicao
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return reply, fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Lendo os status
	log.Printf("Response status: %s", resp.Status)

	var BotkitPayload []BotkitMessage

	body_bytes, _ := io.ReadAll(resp.Body)
	log.Printf("Raw responde from BotKit: %s", string(body_bytes))

	if string(body_bytes) == "" {
		reply = append(reply, "Estamos esperando uma mensagem do BotKit! (Ainda não há uma resposta no fluxo para isso)") // isso com certeza vai mudar
	} else {
		err = json.Unmarshal(body_bytes, &BotkitPayload)
		if err != nil {
			return reply, fmt.Errorf("Error parsing JSON received from BotKit: %v", err)
		}
		// Colocando cada mensgem em um vetor de strings
		for _, msg := range BotkitPayload {
			var innerMsg BotkitInnerMessage
			err := json.Unmarshal([]byte(msg.Text), &innerMsg)
			if err != nil {
				return reply, fmt.Errorf("Error parsing json in internal messages: %v", err)
			}
			reply = append(reply, innerMsg.Body)
		}
	}

	return reply, nil
}
