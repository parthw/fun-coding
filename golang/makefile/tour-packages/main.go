package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1000)
	fmt.Println(rand.Intn(10))
}
