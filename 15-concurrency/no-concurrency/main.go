package main

import "fmt"

func main() {
	hello()
	world()
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
