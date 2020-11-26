package main

import (
	"fmt"
	"math"
)

// Function Values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Function Closures
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	hypot := func (x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(4,5))
	fmt.Println(compute(hypot))

	closureExample := adder()
	for i:=0; i < 10; i++ {
		fmt.Println(closureExample(i))
	}
}