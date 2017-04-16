package main

import "fmt"

func main() {
	ages := make(map[string]int) // mapping from string to int
	ages["alice"] = 31
	ages["thomas"] = 25

	// vs
	hobbies := map[string]string{
		"alice":  "basketball",
		"thomas": "ballerina",
	}

	fmt.Println(ages)
	fmt.Println(hobbies)

	fmt.Println(ages["alice"])
	fmt.Println(hobbies["thomas"])

	fmt.Println(ages["dimas"])    // 0 (default value of int)
	fmt.Println(hobbies["dimas"]) // "" (default value of string)
}
