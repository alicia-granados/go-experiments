package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(localhost:3306)/database
const url = "root:@tcp(localhost:3306)/blog_db"

// giarda la conexion
var db *sql.DB

// realiza coneccion
func Connect() {
	conecton, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Coneccion exitosa")
	db = conecton
}

// cierra la conexion
func Close() {
	db.Close()
}

// verficar la conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
