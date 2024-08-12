package main

import (
	"github.com/PoliGNU-USP/projetoSUS-gateway/routes"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Criando os nossos logs
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("Inicializando o Gateway")

	logger.Info("Pegando variáveis de ambiente")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Criando o router mux para gerenciar os paths http
	router := http.NewServeMux()

	// Dizendo o que acontece quando o nosso router recebe uma requisicao post no endereco padrao
	// no caso, vamos mandar pra essa funcao ReceiveReply
	router.HandleFunc("POST /", routes.ReceiveReply) // aqui que a mágica acontece

	// Iniciando o servidor
	server := http.Server{
		Addr:    os.Getenv("DEV_LOCALHOST"),
		Handler: router,
	}

	logger.Info("Listening na porta " + server.Addr)
	log.Fatal(server.ListenAndServe())
}
