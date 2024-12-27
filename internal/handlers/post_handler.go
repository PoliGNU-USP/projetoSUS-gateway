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
	log.Println("Recebi mensagem do WhatsApp")
	if r.Method != http.MethodPost {
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	// Pegando a mensagem que vem da twilio
	twilio_message, err := utils.ParseTwilioRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Println("A mensagem depois do parser é essa:")
	log.Println(twilio_message)

	// Transformando a mensagem da twilio em uma mensagem para a Conversation
	user_message := models.Message{
		Sender:    twilio_message.AccountSid,
		Text:      twilio_message.Body,
		Timestamp: time.Now(),
	}

	log.Printf("Recebi a mensagem: %v do número: %v \n", user_message.Text, user_message.Sender)

	// Salvando a mensagem na conversa
	log.Println("Salvando a mensagem no banco de dados de conversa")
	repositories.SaveMessage(twilio_message.AccountSid, user_message)

	// Enviar a mensagem pro BotKit
	log.Println("Enviando a mensagem pro BotKit")
	reply, err := services.SendToBotkit(*twilio_message)
	if err != nil {
		log.Printf("Erro mandando mensagem para o BotKit: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("Recebi do BotKit: ", reply)

	// Para cada mensagem do bot, cria uma message e salva no banco de dados
	for _, msg := range reply {
		bot_message := models.Message{
			Sender:    "BotKit",
			Text:      msg,
			Timestamp: time.Now(),
		}

		repositories.SaveMessage(twilio_message.AccountSid, bot_message)
	}

	log.Println("Enviando pro usuário as mensagem", reply)
	if err := services.RespondToUser(w, reply); err != nil {
		log.Printf("Erro enviando a mensagem para o usuário: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
