package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p *Point) Distance(another Point) float64 {
	x := p.x - another.x
	y := p.y - another.y
	return math.Sqrt(x*x + y*y)
}

func main() {
	first := NewPoint(2, 4)
	second := NewPoint(2, 1)

	distance := first.Distance(second)
	fmt.Println(distance)
}
