package main

import (
	"fmt"
)

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{
		sideLength: 10,
	}

	t := triangle{
		base:   10,
		height: 10,
	}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println("The area is", area)
}
