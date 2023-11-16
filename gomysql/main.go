package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	//db.Ping()
	fmt.Println(db.ExistTable("users"))
	//db.CreateTable(models.UserSchema, "users")

	//db.TruncateTable("users")

	user := models.CreateUser("alex", "alex123", "alex@gmail.com")
	fmt.Println(user)
	db.Close()
}
