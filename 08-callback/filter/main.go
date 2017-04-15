package main

import "fmt"

func main() {
	filtered := filter([]int{1, 2, 3, 4, 5}, func(num int) bool {
		return num > 1
	})
	fmt.Println(filtered)

	filtered2 := filter([]int{1, 2, 3, 4, 5}, func(num int) bool {
		return num < 4
	})
	fmt.Println(filtered2)
}

func filter(numbers []int, callback func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if callback(num) {
			result = append(result, num)
		}
	}
	return result
}
