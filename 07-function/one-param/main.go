package main

import "fmt"

func main() {
	greet("budi")
	greet("jonie")
}

func greet(name string) {
	fmt.Println("Hey", name)
}
