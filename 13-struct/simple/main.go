package main

import (
	"fmt"
	"time"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
	Birthday  time.Time
}

func main() {
	var budi Student
	budi.FirstName = "Budi"
	budi.LastName = "Irawan"

	fmt.Println(budi)

	firstName := &budi.FirstName
	*firstName = *firstName + " ganteng"

	fmt.Println(budi)
}
