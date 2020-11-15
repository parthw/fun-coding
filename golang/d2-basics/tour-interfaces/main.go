package main

import (
	"fmt"
	"math"
)

type geometry interface {
	Area() float64
	Perimeter() float64
}

type rect struct {
	l, b float64
}

type circle struct {
	r float64
}

// Area Circle
func (c circle) Area() float64 {
	return math.Pi * c.r * c.r
}

// Perimeter Circle
func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

// Area rect
func (r rect) Area() float64 {
	return r.l * r.b
}

// Perimeter rect
func (r rect) Perimeter() float64 {
	return 2 * r.l + 2 * r.b
}

func main()  {
	r := rect{2,4}
	c := circle{100}

	var g geometry
	g = r
	fmt.Println(g.Area(), g.Perimeter())

	g = c
	fmt.Println(g.Area(), g.Perimeter())
}

