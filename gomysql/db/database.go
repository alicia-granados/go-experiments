package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(localhost:3306)/database

const url = "root:@tcp(localhost:3306)/goweb_db"

var db *sql.DB

// coneccion
func Connect() {
	conecton, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Coneccion exitosa")
	db = conecton
}

func Close() {
	db.Close()
}
