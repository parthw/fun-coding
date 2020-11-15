package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Println("Enter a float digit for sqrt")
	fmt.Scanln(&x)

	fmt.Println(math.Sqrt(x), "👍")

}
