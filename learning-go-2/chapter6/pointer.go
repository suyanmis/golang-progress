package main

type person struct {
	FirstName string
	LastName  *string
	Age       int
}

func main() {
	p := person{
		FirstName: "Pound",
		LastName:  makePointer("Pepe"),
		Age:       25,
	}
}

func makePointer[T any](t T) *T {
	return &t
}
