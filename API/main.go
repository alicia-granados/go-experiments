package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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
