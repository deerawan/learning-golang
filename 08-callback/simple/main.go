package main

import "fmt"

func main() {
	myNumbers := []int{2, 4, 6, 8}
	mapCalculate(myNumbers, func(num int) {
		num = num * 2
		fmt.Println(num)
	})
}

func mapCalculate(numbers []int, callback func(int)) {
	for _, num := range numbers {
		callback(num)
	}
}
