package main

import "fmt"

// Vertex struct
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 0}
	fmt.Println(Vertex{1, 2})
	fmt.Println(v.X, v.Y)

	p := &v
	(*p).X = 1e9
	fmt.Println(v)

	v.X = 10
	fmt.Println(p)
	fmt.Println(v)

	z := Vertex{X: 1}
	fmt.Println(z)
}
