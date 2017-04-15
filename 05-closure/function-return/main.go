package main

import "fmt"

func increment() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func main() {
	myIncrement := increment()
	fmt.Println(myIncrement())
	fmt.Println(myIncrement())
}
