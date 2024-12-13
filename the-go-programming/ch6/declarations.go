package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Path []Point

func (p Point) Distance(point Point) float64 {
	return math.Hypot((p.X - point.X), (p.Y - point.Y))
}
func main() {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{4, 6}
	fmt.Println(p1.Distance(p2))
	path := Path{Point{1, 1}, Point{5, 1}, Point{5, 4}, Point{1, 1}}

	fmt.Printf("path.Distance(): %v\n", path.Distance())
}

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i].Distance(p[i-1])
		}
	}
	return sum
}
