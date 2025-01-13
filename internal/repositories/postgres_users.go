package repositories

import (
	"gateway/internal/models"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao inicializar PostgresDB: %v", err)
	}

	db.AutoMigrate(&models.User{})

	DB = db

	return DB, nil
}

func DisconnectPostgresDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Erro ao obter conexão SQL do GORM: ", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("Erro ao fechar a conexão com o banco de dados: ", err)
	} else {
		log.Println("Conexão com o banco de dados fechada com sucesso.")
	}
}

func SaveUser(userData models.UserData) {
	// Se um usuário é passado dentro de UserData salva esse usuário,
	// se não, salva o nome de usuário e telefone da mensagem por padrão

	var phoneNumber string
	// Apenas armazena o DDD e o número do telefone
	if userData.Msg != nil {
		phoneNumber = strings.TrimPrefix(userData.Msg.From, "whatsapp:+55")
	} else {
		phoneNumber = userData.User.PhoneNumber
	}

	_, err := SearchPhoneNumber(phoneNumber)

	// Usuário já existe ou erro ao verificar usuário
	if err == nil || err != gorm.ErrRecordNotFound {
		return
	}

	if userData.User != (models.User{}) {
		newUser := models.User{
			Name:        userData.User.Name,
			PhoneNumber: userData.User.PhoneNumber,
			Team:        userData.User.Team,
		}

		if err := DB.Create(&newUser).Error; err != nil {
			panic("Falha ao adicionar usuário.")
		}

		return
	}

	// Cria o usuário com o nome e o telefone definidos na mensagem
	newUser := models.User{
		Name:        userData.Msg.ProfileName,
		PhoneNumber: phoneNumber,
	}

	if err := DB.Create(&newUser).Error; err != nil {
		panic("Falha ao adicionar pessoa.")
	}
}

func SearchPhoneNumber(number string) (models.User, error) {
	var user models.User
	// Busca o número no banco de dados
	result := DB.Where("phone_number = ?", number).First(&user)

	// Verifica se o usuário já existe no banco de dados
	if result.Error == nil {
		log.Printf("Usuário já existe no banco de dados: %v", user)
		return user, nil
	}

	if result.Error != gorm.ErrRecordNotFound {
		log.Printf("Erro ao verificar usuário: %v", result.Error)
	}

	return models.User{}, result.Error
}
