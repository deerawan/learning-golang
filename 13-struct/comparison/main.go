package main

import "fmt"

type Point struct{ X, Y int }

func main() {
	pointA := Point{1, 2}
	pointB := Point{1, 2}

	fmt.Println(pointA == pointB)

	points := make(map[Point]int)
	points[Point{3, 4}]++
	points[Point{5, 6}] = 2

	fmt.Println(points)
}
