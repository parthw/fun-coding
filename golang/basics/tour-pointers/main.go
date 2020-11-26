package main

import "fmt"

func main() {
	x := 100
	p := &x

	fmt.Println(p)
	fmt.Println(*p)

	*p = 21
	fmt.Println(x)
	fmt.Println(*p)
	fmt.Println(p)
}
