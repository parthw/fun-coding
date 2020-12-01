package main

// newgrpc is leading to forward compaitability by
// adding struct UnimplementedFuncStruct

import (
	"fmt"

	"example.com/understanding-grpc-change/newgrpc"
	"example.com/understanding-grpc-change/oldgrpc"
)

type mynewgrpc struct {
	newgrpc.UnimplementedFuncStruct
}

// If this gets commented then called from newgrpc package will be the output
func (mynewgrpc) IsThisOverriding() {
	fmt.Println("called from main.go")
}

type myoldgrpc struct{}

func (myoldgrpc) Myfunction() {}

func main() {
	n := &mynewgrpc{}
	newgrpc.Register(n)

	n.IsThisOverriding()

	o := &myoldgrpc{}
	oldgrpc.Register(o)
}
