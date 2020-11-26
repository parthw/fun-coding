package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	if sum < 100 {
		fmt.Printf("%d is less than 100", sum)
	} else {
		fmt.Printf("%d is not less than 100", sum)
	}
}
