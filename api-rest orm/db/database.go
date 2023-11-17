package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
var dsn = "root:@tcp(localhost:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"

var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("error en la conexion", err)
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
		return db
	}
}()

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

// verifica si existe una tabla o no
func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s' ", tableName)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// recorrer la tabla
	return rows.Next()
}

// Crea una tabla
func CreateTable(schema string, name string) {
	if !ExistTable(name) {
		_, err := Exec(schema)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

// reiniciar el registro de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	db.Exec(sql)
}

// NOTA: db.Query o db.Exec ya se pueden usar sin db.tal
// polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
