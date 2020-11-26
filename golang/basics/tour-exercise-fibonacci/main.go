package main

import "fmt"

// With Closure
func fiboClosure() func(int) {
	a := 0
	b := 1
	return func(i int) {
		fmt.Println(a)
		nextTerm := a + b
		a = b
		b = nextTerm
	}
} 

func main() {
	n := 10

	// Closure Example
	fClosure := fiboClosure()
	for i := 0; i < n; i++ {
		fClosure(i)
	}

}