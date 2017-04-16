package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5}

	fmt.Println(len(numbers))
	fmt.Println(numbers[1])
	fmt.Println(numbers[1:4])
	fmt.Println("budi"[1]) // will print ascii decimal of "u"

	anotherNumbers := numbers[2:3]
	fmt.Println("length: ", len(anotherNumbers))   // returns 1 => 3
	fmt.Println("capacity: ", cap(anotherNumbers)) // returns 3 => 3, 4, 5

	superNumbers := numbers[:3] // = [0:3] => 1, 2, 3
	fmt.Println(superNumbers)

	deathNumbers := numbers[2:] // = [2:len()] => 3, 4, 5
	fmt.Println(deathNumbers)

	allNumbers := numbers[:]
	fmt.Println(allNumbers)
}
