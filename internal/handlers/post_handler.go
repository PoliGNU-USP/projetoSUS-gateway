package handlers

import (
	"gateway/internal/services"
	"gateway/internal/utils"
	"log"
	"net/http"
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

	// TODO: Salvar a mensagem do usuário no ChatHistory

	// Enviar a mensagem pro BotKit
	reply, err := services.SendToBotkit(*twilio_message)
	if err != nil {
		log.Printf("Erro mandando mensagem para o BotKit: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	//TODO: Salvar as mensagem do bot no ChatHistory
	//Utilizar os cookies da twilio?
	//https://www.twilio.com/docs/messaging/twiml#cookies

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
