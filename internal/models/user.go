package models

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	PhoneNumber string `gorm:"size:100;not null"`
	Team        string `gorm:"size:100;not null"`
}

type UserData struct {
	User User
	Msg  *TwilioMessage
}
