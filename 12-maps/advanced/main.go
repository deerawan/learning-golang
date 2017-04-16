package main

import "fmt"

func main() {
	ages := make(map[string]int) // mapping from string to int
	ages["alice"] = 31
	ages["thomas"] = 25

	for name, age := range ages {
		fmt.Printf("%s with age %d \n", name, age)
	}

	age, ok := ages["bob"]
	fmt.Println(age) // 0
	fmt.Println(ok)  // false, useful to check in if/else statement
}
