package main

import (
	"fmt"
)

type Employee struct {
	ID        int
	Name      string
	Position  string
	Salary    int
	ManagerID int
}

// IF ALL fields are comparable, then structs can be compared too
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
	p1 := Point{1, 2}
	p2 := Point{2, 1}
	fmt.Println(p1.X == p2.X, p1.Y == p2.Y)
	fmt.Println(p1 == p2)

	// Comparable struct types can be used as key in maps

	pointExists := make(map[Point]bool)

	pointExists[p1] = true
	pointExists[p2] = true
	if p := pointExists[p1]; !p {
		fmt.Println("Creating the point.")
		pointExists[p1] = true
	} else {
		fmt.Println("Already exists.", p)
	}

	var w Wheel
	w.Radius = 5
	w.X = 8 // works only when used embedding. Instead of Circle Circle -> Circle
	w.Y = 8 // equivalent to w.Circle.Point.Y

	fmt.Printf("%#v\n", w)
	// dot format
	var w2 Wheel = Wheel{Circle: Circle{Point: Point{X: 1, Y: 1}, Radius: 5}, Spokes: 5}

	fmt.Println(w2)
}
