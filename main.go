package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)

	router.Run("localhost:8080")
}

// album represents data about a record album.
type pizza struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Size  string  `json:"size"`
	Price float64 `json:"price"`
}

var pizzas = []pizza{
	{ID: "1", Title: "Chocolate", Size: "35", Price: 56.99},
	{ID: "2", Title: "Barbecue", Size: "30", Price: 17.99},
	{ID: "3", Title: "Peperoni", Size: "45", Price: 39.99},
}

// getPizzas responds with the list of all albums as JSON.
func getPizzas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pizzas)
}

// postAlbums adds an album from JSON received in the request body.
func postPizzas(c *gin.Context) {
	var newPizza pizza

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newPizza); err != nil {
		return
	}

	// Add the new album to the slice.
	pizzas = append(pizzas, newPizza)
	c.IndentedJSON(http.StatusCreated, newPizza)
}

// getPizzasByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getPizzasByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range pizzas {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
