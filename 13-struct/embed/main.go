package main

import "fmt"

type Point struct{ X, Y int }

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var wheel Wheel
	wheel.X = 10
	wheel.Y = 15
	wheel.Radius = 20
	wheel.Spokes = 1

	fmt.Println(wheel)

	wheel2 := Wheel{Circle{Point{10, 15}, 20}, 1}

	fmt.Println(wheel2)

	wheel3 := Wheel{
		Circle: Circle{
			Point:  Point{10, 15},
			Radius: 20,
		},
		Spokes: 1,
	}

	fmt.Println(wheel3)
	fmt.Printf("%#v\n", wheel3)
}
