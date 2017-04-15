package main

import (
	"fmt"
	"learning-golang/02-package/calculator"
)

func main() {
	fmt.Println("hello package")
	fmt.Println(calculator.Brand)
	fmt.Println(calculator.Add(1, 2))
	fmt.Println(calculator.Multiply(10, 20))
}
