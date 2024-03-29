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

	//users := models.ListUser()

	user := models.GetUser(2)
	fmt.Println(user)

	user.Username = "Juan"
	user.Email = "juan@gmail.com"
	user.Save()
	fmt.Println(models.ListUser())

	user2 := models.GetUser(1)
	fmt.Println(user2)
	user2.Delete()
	fmt.Println(models.ListUser())

	//db.TruncateTable("users")
	fmt.Println(models.ListUser())

	db.Close()
}
