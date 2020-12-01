package newgrpc

import "fmt"

// NewInt interface
type NewInt interface {
	// Here Something is not implemented in main.go leading to forward compaitibility
	Something()
	// To check whether IsThisOverriding() gets overide by main.go or not
	IsThisOverriding()
	myfunction()
}

// UnimplementedFuncStruct struct
type UnimplementedFuncStruct struct{}

// Register function
func Register(n NewInt) {}

func (UnimplementedFuncStruct) myfunction() {}

// Something is not implemented in main.go leading to forward compaitibility
func (UnimplementedFuncStruct) Something() {}

// IsThisOverriding getting overide by main.go or not
func (UnimplementedFuncStruct) IsThisOverriding() {
	fmt.Println("called from newgrpc package")
}
