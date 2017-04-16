package main

import "fmt"
import "encoding/json"
import "log"

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Color  bool     `json:"color"`
	Actors []string `json:"actors"`
	Grammy bool     `json:"grammy"`
}

func main() {
	var movies = []Movie{
		{
			Title:  "Casablanca",
			Year:   1942,
			Color:  false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
			Grammy: true,
		},
		{
			Title:  "Hand Luke",
			Year:   1967,
			Color:  true,
			Actors: []string{"Steve McQueen"},
			Grammy: false,
		},
	}

	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Printf("%T\n", data)
	fmt.Println(data)
}
