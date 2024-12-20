package services

import (
	"encoding/xml"
	"gateway/internal/models"
	"net/http"
)

func RespondToUser(w http.ResponseWriter, messages []string) error {

	var twimlMessages []models.TwiML_Message

	// Adiciono cada mensagem no vetor
	for _, msg := range messages {
		twimlMessages = append(twimlMessages, models.TwiML_Message{Body: msg})
	}

	// Transformo o vetor no formato esperado por eles
	response := models.TwiML{Messages: twimlMessages}

	// Crio os bytes que serão enviados
	xmlbytes, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write(xmlbytes)

	return nil
}
