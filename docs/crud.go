package main

import (
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
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John",
		Price:  56.99,
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry",
		Price:  65.20,
	},
	{
		ID:     "3",
		Title:  "Joker",
		Artist: "Sunil",
		Price:  65.26,
	},
}

// Handler to return all albums
func getAlbum(c *gin.Context) {
	
	c.IndentedJSON(http.StatusOK, albums)
}

// Handler to add a new album
func postAlbum(c *gin.Context) {
	var newAlbum album

	// Bind the JSON data to the album struct
	if err := c.BindJSON(&newAlbum); err != nil {
		// Respond with a 400 status code if there is an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Add the new album to the albums slice
	albums = append(albums, newAlbum)

	// Respond with the new album data and 201 Created status code
	c.JSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context){
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()

	// Register routes
	router.GET("/albums", getAlbum)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumById)

	// Start the server
	router.Run("localhost:8080")
}
