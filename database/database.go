package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Essa structure é a versão em Go do banco de dados
type User struct {
	gorm.Model
	Name    string `json:"name"`
	Number  int    `gorm:"unique"` // Um número único por usuário
	Address string `json:"address"`
	// TODO: Definir mais parâmetros para o banco de dados
}

func ConnectDatabase() (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})

}
