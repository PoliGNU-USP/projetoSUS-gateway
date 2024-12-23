package services

import (
	"bytes"
	"encoding/json"
	"gateway/internal/config"
	"gateway/internal/models"
	"gateway/internal/utils"
	"net/http"
)

func SendToBotkit(msg models.TwilioMessage) ([]string, error) {

	url := config.Env.BOTKIT_URL

	var reply []string

	botkitPayload := map[string]string{
		"type":    "message",
		"text":    msg.Body,
		"channel": "webhook",
		"user":    msg.ProfileName,
	}

	payload, err := json.Marshal(botkitPayload)
	if err != nil {
		return reply, err
	}

	// Criando uma requicao POST com o Json de Payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return reply, err
	}

	// Definindo o tipo de conteudo para application/json
	req.Header.Set("Content-Type", "application/json")

	// Enviando a requisicao
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return reply, err
	}
	defer resp.Body.Close()

	// // Lendo os status
	// log.Printf("Response status: %s", resp.Status)

	reply, err = utils.Botkit_Parser(resp)
	if err != nil {
		return reply, err
	}

	return reply, err

}
