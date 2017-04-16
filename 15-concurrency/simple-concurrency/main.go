package main

import "fmt"

func main() {
	go hello()
	go world()
	// NOTE: the program will print nothing because the main go-routine already finished
	// without waiting for other two go routines
	// Solution: wait group
}

func hello() {
	for i := 0; i < 45; i++ {
		fmt.Println("hello: ", i)
	}
}

func world() {
	for i := 0; i < 45; i++ {
		fmt.Println("world: ", i)
	}
}
