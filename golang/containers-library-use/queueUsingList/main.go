package main

import "container/list"

//Person struct
type Person struct {
	name string
	age  int
}

var q *list.List

func initQueue() *list.List {
	return q.Init()
}

func push(p *Person) {
	q.PushFront(p)
}

func poll() *Person {
	return q.Remove(q.Back()).(*Person)
}

func main() {

}
