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

	//user := models.CreateUser("alexs", "alex1234", "alex@gmail.com")

	users := models.ListUser()
	fmt.Println(users)

	db.Close()
}
