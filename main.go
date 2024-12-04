package main

import (
	"net/http"           // Provides HTTP server and client functionality.
	"github.com/gin-gonic/gin" // Gin framework for building APIs.
)

// album represents data for an album, with JSON tags for serialization.
type album struct {
	ID     string  `json:"id"`     // Album ID
	Title  string  `json:"title"`  // Album title
	Artist string  `json:"artist"` // Artist name
	Price  float64 `json:"price"`  // Album price
}

// Sample data for albums.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// homeHandler responds with a welcome message.
func homeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the Albums API! Use /albums to view the list of albums.")
}

func main() {
	router := gin.Default()                 // Initialize the router.
	router.GET("/", homeHandler)            // Route for the root URL.
	router.GET("/albums", getAlbums)        // Route to get all albums.
	router.Run("localhost:8080")            // Start the server on localhost:8080.
}
