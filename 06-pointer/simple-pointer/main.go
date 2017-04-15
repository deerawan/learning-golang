package main

import "fmt"

func main() {
	x := 10
	fmt.Println("x", x)
	fmt.Println("x address", &x)

	// y type is *int and it points to memory address of x
	var y *int = &x
	fmt.Println("y address", y)
	fmt.Println("y", *y)
}
