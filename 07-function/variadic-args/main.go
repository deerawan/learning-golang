package main

import "fmt"

func main() {
	numbers := []float64{1, 2, 3, 4, 5}
	total := add(numbers...) // numbers is one data that contain slice
	fmt.Println(total)
}

func add(numbers ...float64) float64 {
	fmt.Printf("%T \n", numbers)
	var total float64
	for _, number := range numbers {
		total += number
	}

	return total
}
