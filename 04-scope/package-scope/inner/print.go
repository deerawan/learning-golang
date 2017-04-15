package inner

import "fmt"

// Print everything that we have
func PrintEverything() {
	// initialNumber can't be accessed even in same package name
	fmt.Println(initialNumber)
}
