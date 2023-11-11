package main

import (
	"fmt"
	"net/http"

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

func getAlbums(c *gin.Context) {
	//SERIALIZAR
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	// recibir una estructura -> cargar datos a la esctructura
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	//agregar al album
	albums = append(albums, newAlbum)
	//responder al cliente
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func main() {
	fmt.Println("Bienvenido a la api")

	router := gin.Default()
	/*
		router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hola mundo",
			})
		})*/

	router.GET("/albums", getAlbums)

	router.POST("/albums", postAlbums)

	router.Run("localhost:8081")
}
