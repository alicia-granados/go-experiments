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
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

//listar todos los registros

func ListUser() (Users, error) {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	if rows, err := db.Query(sql); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			user := User{}
			rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
			users = append(users, user)
		}
		return users, nil
	}
}

// obtener un registro
func GetUser(id int) (*User, error) {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id = ?"
	if rows, err := db.Query(sql, id); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		}
		return user, nil
	}
}

// actualizar registro privado
func (user *User) update() {
	sql := "UPDATE users SET  username=?, password= ? , email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

// guardar o editar registro - publico
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

// eliminar un registro
func (user *User) Delete() {
	sql := "DELETE FROM users  WHERE id=?"
	db.Exec(sql, user.Id)
}
