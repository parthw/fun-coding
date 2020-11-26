package main

import "fmt"

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}

func main() {
	s := []int{1,2,3,7,1,9}
	ch := make(chan int)
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)
	x := <- ch
	fmt.Println(x)
	fmt.Println(<- ch)
}