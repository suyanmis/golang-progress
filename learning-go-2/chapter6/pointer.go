package main

import "fmt"

type person struct {
	FirstName string
	LastName  *string
	Age       int
}

func main() {
	p := person{
		FirstName: "Pound",
		LastName:  makePointer("Pepe"), // this works because when a constant passes to a function it is copied to the param(here it is t). Unlike primitives, variables have address in the memory.
		Age:       25,
	}

	fmt.Println(p)

	// while working on primitive type, we can't address the type directly. We need to initiate value first.
	address := 15
	testValue := &address
	fmt.Println(testValue)

	// on the other hand struct types can be referenced immediately.
	p1 := &person{FirstName: "Test"}
	// or
	p2 := &person{}
	fmt.Println(*p1)
	fmt.Println(*p2)

	// null pointer dereferencing
	var a2 *int
	fmt.Println(a2 == nil)
	// this will panic and leave. We have to check null pointer that we want to dereference.
	// fmt.Println(*a2)

}

func makePointer[T any](t T) *T {
	return &t
}
