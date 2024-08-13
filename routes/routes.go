package routes

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PoliGNU-USP/projetoSUS-gateway/communication"
)

func ReceiveReply(w http.ResponseWriter, r *http.Request) {

	//Respondendo rapidamente com um status ok
	w.WriteHeader(http.StatusOK)

	// Pegando as variáveis de ambiente necessárias
	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("Mensagem recebida!")

	// Pegando o corpo da mensagem recebida
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error getting the body: ", err)
	}

	// Pegando as informações do body
	values, err := url.ParseQuery(string(body))
	if err != nil {
		log.Fatal("Error parsing the query: ", err)
	}

	// Construindo a mensagem recebida com algumas informações
	message := communication.Message{
		AccountSid: values.Get("AccountSid"),
		Body:       values.Get("Body"),
		From:       values.Get("From"),
		To:         values.Get("To"),
		// Há mais informações dentro de values que poderiamos pegar
	}

	// Logging da mensagem para vermos
	logger.Info("Mensagem recebida foi: " + message.Body)

	// Usando paralelismo aqui
	// Enviando a mensagem para o BotKit e guardando a resposta
	go func() {
		reply, err := communication.SendToBotKit(message)
		if err != nil {
			fmt.Println(err)
		}

		logger.Info("Esperando resposta do BotKit")

		// URL para o qual mandaremos a mensagem de volta ao usuário
		urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSID + "/Messages.json"

		// reply pode ser um vetor de mensagem
		for _, msg := range reply {
			logger.Info("Enviando mensagem: " + msg)

			msgData := url.Values{}

			// Envio a mensagem para quem eu recebi
			msgData.Set("From", message.To) // No futuro aqui sera o numero do botkit
			msgData.Set("To", message.From)
			msgData.Set("Body", msg) // Conteudo da mensagem

			// Codificando a mensagem que sera enviada
			msgDataReader := *strings.NewReader(msgData.Encode())

			// Criando a requisicao http
			req, err := http.NewRequest("POST", urlStr, &msgDataReader)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.SetBasicAuth(accountSID, authToken)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			// Enviando a mesagem ao Twilio
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			logger.Info("Status do envio ao Twilio: " + res.Status)

			defer res.Body.Close()
		}

	}()

}
