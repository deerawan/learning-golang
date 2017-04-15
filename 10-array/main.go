package main

import "fmt"

func main() {
	// if without number inside brackets e.g []string, it means "slice"
	// array in GO is not dynamic, slice is
	var names [5]string
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(names[1])
	// fmt.Println(names[8]) => will print error out of bounds

	var numbers [5]int
	fmt.Println(numbers) // will print 5 zeroes
}
