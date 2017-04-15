package main

import "fmt"

func main() {

	// func greeting() => we can't do this in golang, in JS we can :)

	greeting := func() {
		fmt.Println("hello world")
	}

	greeting()
}
