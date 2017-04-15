package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
}

func printMessage() {
	// can't access x because x is in  block scope of main()
	fmt.Println(x)
}
