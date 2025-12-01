package main

import (
	"api-pizza/models"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.GET("/pizzas/:id", getPizzasByID)
	router.POST("/pizzas", postPizzas)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

// album represents data about a record album.

var pizzas []models.Pizza

// getPizzas responds with the list of all albums as JSON.
func getPizzas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pizzas)
}

// postAlbums adds an album from JSON received in the request body.
func postPizzas(c *gin.Context) {
	var newPizza models.Pizza

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newPizza); err != nil {
		return
	}
	newPizza.ID = strconv.Itoa(1 + len(pizzas))
	// Add the new album to the slice.
	pizzas = append(pizzas, newPizza)
	savePizza()
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

func loadPizzas() {
	file, err := os.Open("data/pizza.json")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		panic(err)
	}
}

func savePizza() {
	file, err := os.Create("data/pizza.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		panic(err)
	}

}
