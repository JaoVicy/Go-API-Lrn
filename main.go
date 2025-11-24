package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func main() {
	router := gin.Default()
	router.GET("/albums", getPizzas)

	router.Run("localhost:8080")
}
