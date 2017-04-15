package main

import "fmt"

func main() {
	numbers := add(1, 2, 3, 10, 20, 30)
	fmt.Println(numbers)
}

// variadic here same like javascript rest parameters
// numbers type is []float64 => slice (not array)
func add(numbers ...float64) float64 {
	fmt.Printf("%T \n", numbers)
	var total float64
	for _, number := range numbers {
		total += number
	}

	return total
}
