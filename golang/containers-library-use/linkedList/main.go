package main

import (
	"container/list"
	"fmt"
)

// Person struct
type Person struct {
	name string
	age  int
}

func main() {
	l := list.New()
	e10 := l.PushBack(&Person{name: "ten", age: 10})
	e9 := l.PushFront(&Person{name: "nine", age: 9})
	fmt.Println(e9.Value, e10.Value)
	for x := l.Front(); x != nil; x = x.Next() {
		t := x.Value
		fmt.Println(t.(*Person).age)
	}
	lastE := l.Remove(l.Back())
	if lastE == e10.Value {
		fmt.Println("Last Element e10 removed - ", lastE)
	}
}
