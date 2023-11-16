package models

import "gomysql/db"

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

const UserSchema string = `CREATE TABLE users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

// construir ususario
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

// Crear usuario e insertar a db
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.insert()
	return user
}

// insertar registro metodo privado
func (user *User) insert() {
	sql := "INSERT users SET username=?, password= ?, email=?"
	db.Exec(sql, user.Username, user.Password, user.Email)
}
