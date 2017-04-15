package main

import "fmt"

func change(z int) {
	fmt.Println("z addr", &z)
	z = 0
}

func main() {
	x := 5
	fmt.Println("x addr", &x)
	change(x) // try to change x into 0

	fmt.Println("x", x) // x is still 5
}
