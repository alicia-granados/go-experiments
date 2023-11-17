package models

import (
	"apirest-gorm/db"
)

type User struct {
	Id       int64  `json:"id"`       // `xml:  "id"` Para responder con xml
	Username string `json:"username"` // `xml:  "username"` Para responder con yaml
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

func MigrarUser() {
	//migrar una estructura
	db.Database.AutoMigrate(User{})
}
