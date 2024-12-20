package config

import (
	"os"
)

type Config struct {
	Port        string
	PostgresDSN string
	MongoDBURI  string
	BotkitURL   string
	TwilioSID   string
}

var Env *Config

func Load() {
	Env = &Config{
		Port:        getEnv("PORT", "8080"),
		PostgresDSN: getEnv("POSTGRES_DSN", "postgres://user:pass@localhost:5432/db"),
		MongoDBURI:  getEnv("MONGO_URI", "mongodb://root:example@mongo:27017/"),
		BotkitURL:   getEnv("BOTKIT_URL", "http://fluxo:3000/api/messages"),
		TwilioSID:   getEnv("TWILIO_SID", "XXXXXXXX"),
	}
}

// Tenta pegar as variaveis de ambiente, se nao encontrar cai no callback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
