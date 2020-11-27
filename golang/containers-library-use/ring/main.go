package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(3)
	r.Value = "Hello"
	r.Next().Value = "How"
	r.Next().Next().Value = "Are"
	r.Next().Next().Next().Value = "You"

	r.Do(func(x interface{}) {
		fmt.Println(x.(string))
	})

	// Create two rings, r1 and r2, of size 3
	r1 := ring.New(3)
	r2 := ring.New(3)

	// Get the length of the ring
	lr := r1.Len()
	ls := r2.Len()

	// Initialize r1 with "GFG"
	for i := 0; i < lr-1; i++ {
		r1.Value = "GFG"
		r1 = r1.Next()
	}

	// Initialize r2 with "GOLang"
	for j := 0; j < ls; j++ {
		r2.Value = "GoLang"
		r2 = r2.Next()
	}

	// Link ring r1 and ring r2
	rs := r1.Link(r2)

	// Iterate through the combined
	// ring and print its contents
	rs.Do(func(p interface{}) {
		fmt.Println(p.(string))
	})
}
