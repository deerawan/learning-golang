package main

import "fmt"

// variable is delcared in package scope
var message = "hello"

func main() {
	fmt.Println(message)
	printMsg()
	printNumber() // still can print the value 100
}

func printMsg() {
	fmt.Println(message)
}

func printNumber() {
	fmt.Println(number)
}

// even though we put number below main() where it is called, go can still detect it
var number = 100
