package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
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

func getAlbumsID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func deleteAlbumID(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for indx, a := range albums {
		if a.ID == id {
			index = indx
			//c.IndentedJSON(http.StatusOK, a)
			albums = append(albums[:index], albums[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album eliminado"})
			//return
		}
	}
	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
	}

}

func patchAlbumID(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for indx, a := range albums {
		if a.ID == id {
			index = indx

			var newAlbum album
			// recibir una estructura -> cargar datos a la esctructura
			if err := c.BindJSON(&newAlbum); err != nil {
				return
			}

			albums[indx] = newAlbum
			//c.IndentedJSON(http.StatusOK, a)

			c.IndentedJSON(http.StatusOK, gin.H{"message": "album actualizado"})
			//return
		}
	}
	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
	}

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
	router.GET("/albums/:id", getAlbumsID)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbumID)
	router.PATCH("/albums/:id", patchAlbumID)

	router.Run("localhost:8081")
}
