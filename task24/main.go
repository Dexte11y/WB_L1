// Разработать программу нахождения расстояния между
// двумя точками, которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func getDistance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := newPoint(0, 0)
	p2 := newPoint(2, 2)
	fmt.Println(getDistance(*p1, *p2))
}
