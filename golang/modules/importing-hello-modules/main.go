package main

import (
	"fmt"

	"example.com/hello"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	hello.Hello()
	fmt.Println(quoteV3.HelloV3())
}
