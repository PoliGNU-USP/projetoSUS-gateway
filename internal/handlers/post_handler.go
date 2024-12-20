package handlers

import (
	"gateway/internal/models"
	"gateway/internal/repositories"
	"gateway/internal/services"
	"gateway/internal/utils"
	"log"
	"net/http"
	"time"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	// Pegando a mensagem que vem da twilio
	twilio_message, err := utils.ParseTwilioRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Transformando a mensagem da twilio em uma mensagem para a Conversation
	user_message := models.Message{
		Sender:    twilio_message.AccountSid,
		Text:      twilio_message.Body,
		Timestamp: time.Now(),
	}

	// Salvando a mensgem na conversa
	repositories.SaveMessage(twilio_message.AccountSid, user_message)

	// Enviar a mensagem pro BotKit
	reply, err := services.SendToBotkit(*twilio_message)
	if err != nil {
		log.Printf("Erro mandando mensagem para o BotKit: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Para cada mensagem do bot, cria uma message e salva no banco de dados
	for _, msg := range reply {
		bot_message := models.Message{
			Sender:    "BotKit",
			Text:      msg,
			Timestamp: time.Now(),
		}

		repositories.SaveMessage(twilio_message.AccountSid, bot_message)
	}

	if err := services.RespondToUser(w, reply); err != nil {
		log.Printf("Erro enviando a mensagem para o usuário: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Save to databases TUDO ERRADO AQUI?
	// if err := services.SaveMessage(twilio_message); err != nil {
	// 	log.Printf("Error saving message: %v", err)
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }
}
