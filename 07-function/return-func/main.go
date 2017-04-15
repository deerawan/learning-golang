package main

import "fmt"

func main() {
	greeting := makeGreeting()
	fmt.Println(greeting())
}

func makeGreeting() func() string {
	return func() string {
		return "hello world"
	}
}
