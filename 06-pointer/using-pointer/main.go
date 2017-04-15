package main

import "fmt"

func change(z *int) { // z store the pointer
	fmt.Println("z addr", z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println("x addr", &x)
	change(&x) // we pass the memory address here

	fmt.Println("x", x)
}
