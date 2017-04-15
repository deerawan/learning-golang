package main

import "fmt"

func main() {
	fmt.Println(greet("John ", "Doe"))
}

func greet(firstName, lastName string) (string, string) {
	return fmt.Sprint(firstName, lastName), fmt.Sprint(firstName, lastName)
}
