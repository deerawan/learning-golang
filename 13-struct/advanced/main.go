package main

import "fmt"

type Student struct {
	ID                  int
	FirstName, LastName string
	Address             string
}

func main() {
	// how to declare with values
	tono := Student{
		FirstName: "tono",
		LastName:  "irawan",
	}
	fmt.Println(tono) // ID is 0

	assignID(&tono)

	fmt.Println(tono) // ID is 1
}

func assignID(student *Student) {
	student.ID = 1
}
