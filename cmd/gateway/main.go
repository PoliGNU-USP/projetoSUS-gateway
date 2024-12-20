package main

import (
	"gateway/internal/config"
	"gateway/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Carregando as configuracoes
	config.Load()

	// Inicializando o router
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handlers.HandlePost)

	// Iniciando o server
	log.Printf("Inicializando o server na porta %s...", config.Env.Port)
	if err := http.ListenAndServe(":"+config.Env.Port, mux); err != nil {
		log.Fatalf("Erro ao inicializar o gateway: %v", err)
	}
}
