package main

import "fmt"

func main() {
	msg := greet("Lara ", "Croft")
	fmt.Println(msg)
}

func greet(firstName, lastName string) string {
	return fmt.Sprint(firstName, lastName)
}
