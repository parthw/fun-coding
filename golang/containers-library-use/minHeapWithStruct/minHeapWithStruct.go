package main

import (
	"container/heap"
	"fmt"
)

// Person struct
type Person struct {
	age  int
	name string
}

// Persons struct
//type Persons []Person
type Persons []*Person

//Len of Total Person
func (p Persons) Len() int { return len(p) }

//Less of Persons age
func (p Persons) Less(i, j int) bool { return p[i].age < p[j].age }

//Swap of Person
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Pop of heap
func (p *Persons) Pop() (lastElement interface{}) {
	old := *p
	lastElementIndex := len(*p) - 1
	lastElement = (*p)[lastElementIndex]
	*p = old[:lastElementIndex]
	return lastElement
}

// Push of heap
func (p *Persons) Push(newElement interface{}) {
	//*p = append(*p, newElement.(Person))
	x := newElement.(*Person)
	*p = append(*p, x)
}

func main() {

	p := &Persons{
		&Person{age: 23, name: "twenty-three"},
		&Person{age: 22, name: "twnety-two"},
		&Person{age: 25, name: "twenty-five"},
		&Person{age: 27, name: "twnety-seven"},
	}
	heap.Init(p)
	heap.Push(p, &Person{age: 10, name: "Ten"})
	fmt.Println("Printing p[0] ", (*(*p)[0]))
	for p.Len() > 0 {
		fmt.Printf("%v ", heap.Pop(p))
	}
}
