package main

import "fmt"

func main() {
	greet("Budi", "Irawan")
}

func greet(firstName, lastName string) {
	fmt.Println("Hey", firstName, lastName)
}
