package main

import (
	"gateway/internal/config"
	"gateway/internal/handlers"
	"gateway/internal/repositories"
	"log"
	"net/http"
)

func main() {
	// Carregando as configuracoes
	config.Load()

	//TODO: As conexões com outros bancos de dados devem ser feitas no começo do código
	//Iniciando a conexão com o BD
	mongoClient, err := repositories.InitMongoDB(config.Env.MongoDBURI, "mydatabase", "conversations")
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(nil)

	// Inicializando o router
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handlers.HandlePost)

	// Iniciando o server
	log.Printf("Inicializando o server na porta %s...", config.Env.Port)
	if err := http.ListenAndServe(":"+config.Env.Port, mux); err != nil {
		log.Fatalf("Erro ao inicializar o gateway: %v", err)
	}
}
