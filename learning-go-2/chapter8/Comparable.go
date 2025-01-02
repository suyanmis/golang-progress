package main

import (
	"cmp"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func OrderPeople(p1, p2 Person) int {
	out := cmp.Compare(p1.Name, p2.Name)
	if out == 0 {
		out = cmp.Compare(p1.Age, p2.Age)
	}
	return out
}

func main() {
	p1 := Person{Name: "Serhat", Age: 28}
	p2 := Person{Name: "Serhat", Age: 30}
	p3 := Person{Name: "Serdar", Age: 32}

	tree := NewTree(OrderPeople)
	tree.Add(p1)
	tree.Add(p2)
	tree.Add(p3)
	l := tree.Traversal()
	fmt.Println(l)
	fmt.Printf("tree.Contains(p3): %v\n", tree.Contains(p3))
}
