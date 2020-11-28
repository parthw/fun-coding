package main

import (
	"fmt"
	"runtime"
	"time"
)

func fetchNumber(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func sum(ch chan int) {
	sum := 0
	for i := range ch {
		sum += i
	}
	fmt.Println(sum)
}

func sumWithInChannels() {

	ch := make(chan int, 10)
	fmt.Println("Print something")
	go fetchNumber(ch)
	go sum(ch)
}

func main() {
	go sumWithInChannels()
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
	time.Sleep(10 * time.Second)
}
