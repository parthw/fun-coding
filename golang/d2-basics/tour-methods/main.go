package main

import (
	"fmt"
	"math"
)

// Vertex struct
type Vertex struct {
	X, Y float64
}

// Abs function same as func Abs(v Vertex) float64
func (v Vertex) Abs() float64 {
	v.X = 5
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// AnotherAbs with pointers
func (v *Vertex) AnotherAbs() float64 {
	v.X = 5
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3,4}
	// Same as fmt.Println(Abs(v))
	fmt.Println(v.Abs())
	fmt.Println(v)

	z := &Vertex{3,4}
	fmt.Println(z.AnotherAbs())
	fmt.Println(*z)

}