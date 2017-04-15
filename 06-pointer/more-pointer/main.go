package main

import "fmt"

func main() {
	x := 42
	y := &x // store x's memory address to y

	fmt.Println("x", x)

	*y = 2 // change the value for this memory address

	fmt.Println("x", x)
}
