package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Name    string
	Age     int
	Parents []string
}

func main() {
	data := []byte(`{"name":"Wednesday","address":"sydney","age":6,"parents":["Gomez","Morticia"]}`)
	var myMessage Message
	err := json.Unmarshal(data, &myMessage)
	if err != nil {
		log.Fatalf("JSON unmarshal failed: %s", err)
	}
	fmt.Printf("%#v", myMessage)

}
