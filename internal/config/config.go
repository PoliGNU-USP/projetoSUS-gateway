package config

import (
	"os"
)

type Config struct {
	Port               string
	POSTGRESS_DSN      string
	MONGODB_URI        string
	BOTKIT_URL         string
	TWILIO_SID         string
	MONGODB_DBNAME     string
	MONGODB_COLLECTION string
}

var Env *Config

func Load() {
	Env = &Config{
		Port:               getEnv("PORT", "8080"),
		POSTGRESS_DSN:      getEnv("POSTGRES_DSN", "postgres://user:pass@localhost:5432/db"),
		MONGODB_DBNAME:     getEnv("MONGODB_DBNAME", "my_database"),
		MONGODB_COLLECTION: getEnv("MONGODB_COLLECTION", "conversations"),
		MONGODB_URI:        getEnv("MONGO_URI", "mongodb://root:example@mongo:27017/"),
		BOTKIT_URL:         getEnv("BOTKIT_URL", "http://fluxo:3000/api/messages"),
		TWILIO_SID:         getEnv("TWILIO_SID", "XXXXXXXX"),
	}
}

// Tenta pegar as variaveis de ambiente, se nao encontrar cai no callback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
