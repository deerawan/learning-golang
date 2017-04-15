package main

import "fmt"

func main() {
	// without defer will print "world hello"
	// world()
	// hello()

	// will print "hello world"
	defer world()
	hello()

	// will print "hello world good morning"
	// First in Last out of defer
	// defer morning()
	// defer world()
	// hello()
}

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func morning() {
	fmt.Println("good morning")
}
