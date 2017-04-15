package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)

	// define inner block scope
	{
		fmt.Println(x)
		y := 50
		fmt.Println(y)
	}

	// fmt.Println(y) => won't be detected
}
