package main

import "fmt"

// Struct Vertex
type Vertex struct {
	Lat, Long float64
}


func main() {
	// Array implementation
	var x [2]int
	x[0], x[1] = 1, 2
	fmt.Println(x)

	// Slice Implementation
	primes := []int{2,3,5,7}
	fmt.Println(len(primes), cap(primes), primes)

	myslice := make([]int, 5)
	fmt.Println(myslice)

	newSlice := append(myslice, 100)
	fmt.Println(newSlice)

	for i, v := range newSlice {
		fmt.Println(i, v)
	}

	// Maps Implementation
	m := make(map[string]Vertex)
	m["India"] = Vertex{100.34324, 101.1231}
	fmt.Println(m, m["India"])

	m2 := map[string]Vertex{
		"India": {100.00, 101.121},
	}
	fmt.Println(m2["India"])
}
