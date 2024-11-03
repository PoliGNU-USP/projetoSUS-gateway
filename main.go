package main

import (
	"github.com/PoliGNU-USP/projetoSUS-gateway/database"
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

	logger.Info("Carregando variáveis de ambiente")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro pegando as variáveis de ambiente", err)
	}

	// Conectando com o banco de dados
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Erro conectando com o banco de dados", err)
	}
	logger.Info("Banco de dados conectado com sucesso!")

	// Criando ou conectando a tabela dos usuarios
	err = db.AutoMigrate(&database.User{})
	if err != nil {
		log.Fatal("Erro ao criar tabela dos usuários", err)
	}

	// Criando o router mux para gerenciar os paths http
	router := http.NewServeMux()

	// Dizendo o que acontece quando o nosso router recebe uma requisicao post no endereco padrao
	// no caso, vamos mandar pra essa funcao ReceiveReply
	router.HandleFunc("POST /", routes.ReceiveReply(db, logger)) // aqui que a mágica acontece

	// Iniciando o servidor
	server := http.Server{
		Addr:    os.Getenv("DEV_LOCALHOST"),
		Handler: router,
	}

	logger.Info("Listening na porta " + server.Addr)
	log.Fatal(server.ListenAndServe())
}
