package main

import "fmt"

// const x string; => will emit error, const need to have its value declared

func main() {
	// const can't be re-assigned
	const y = "world"
	fmt.Println(y)
}
