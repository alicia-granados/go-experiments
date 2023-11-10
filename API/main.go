package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:id`
	Title  string  `json:title`
	Artist string  `json:artist`
	Price  float64 `json:price`
}

var albums = []album{
	{ID: "1", Title: "blue train", Artist: "Jhon", Price: 56.99},
	{ID: "2", Title: "jeru", Artist: "Gerry", Price: 56.99},
	{ID: "3", Title: "sarah vaughan", Artist: "Sarah", Price: 79.99},
}

func main() {
	fmt.Println("Bienvenido a la api")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola mundo",
		})
	})

	router.Run("localhost:8080")
}
