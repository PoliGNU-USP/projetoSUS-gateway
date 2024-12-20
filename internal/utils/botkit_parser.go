package utils

import (
	"encoding/json"
	"gateway/internal/models"
	"io"
	"net/http"
)

func Botkit_Parser(resp *http.Response) ([]string, error) {

	var reply []string

	var BotkitWrapped []models.BotkitWrapper

	body_bytes, _ := io.ReadAll(resp.Body)

	if string(body_bytes) == "" {
		reply = append(reply, "Estamos esperando uma mensagem do BotKit! (Ainda não há uma resposta no fluxo para isso)") // isso com certeza vai mudar
	} else {
		if err := json.Unmarshal(body_bytes, &BotkitWrapped); err != nil {
			return reply, err
		}

		// Colocando cada mensagem no vetor
		for _, msg := range BotkitWrapped {

			var message models.BotkitMessage

			if err := json.Unmarshal([]byte(msg.Text), &message); err != nil {
				return reply, err
			}

			reply = append(reply, message.Body)
		}
	}

	return reply, nil
}
