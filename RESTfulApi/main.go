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

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 36.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Toxicity", Artist: "System of a Down", Price: 54.99},
	{ID: "5", Title: "Yellow submarine", Artist: "The Beatles", Price: 24.90},
	{ID: "6", Title: "Death Magnetic", Artist: "Metallica", Price: 46.90},
}

// function to getAlbums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// function to postAlbums
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to new Album.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new Album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAalbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response
func getAalbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over the list of albums, looking for an album
	// whose ID value matches the parameter.
	for _, v := range albums {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAalbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
