package main

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

func main() {

}
