package main

// album represents data about a record album.
type pizza struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

var pizzas = []pizza{
	{ID: "1", Title: "Blue Train", Type: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Type: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Type: "Sarah Vaughan", Price: 39.99},
}

func main() {

}
