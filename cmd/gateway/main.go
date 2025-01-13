package main

import (
	"log"
	"net/http"

	"gateway/internal/config"
	"gateway/internal/handlers"
	"gateway/internal/repositories"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	log.Println("Iniciando o Gateway")
	// Carregando as configuracoes
	log.Println("Carregando as configurações")
	config.Load()

	// TODO: As conexões com outros bancos de dados devem ser feitas no começo do código
	// Iniciando a conexão com o BD
	log.Println("Iniciando a conexão com o BD Mongo")

	mongoClient, err := repositories.InitMongoDB(config.Env.MONGODB_URI, config.Env.MONGODB_DBNAME, config.Env.MONGODB_COLLECTION)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(nil)

	// PostgresDB
	dsn := "host=postgres user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	dbPostgres, err := repositories.InitPostgresDB(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize PostgresDB: %v", err)
	}
	defer repositories.DisconnectPostgresDB(dbPostgres)

	// Inicializando o router
	log.Println("Iniciando o router MUX")
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handlers.HandlePost)

	// Iniciando o server
	log.Printf("Inicializando o server na porta %s...", config.Env.Port)
	if err := http.ListenAndServe(":"+config.Env.Port, mux); err != nil {
		log.Fatalf("Erro ao inicializar o gateway: %v", err)
	}
}
