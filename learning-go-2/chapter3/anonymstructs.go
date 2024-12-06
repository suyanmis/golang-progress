package main

import "fmt"

var person1 struct {
	name string
	age  int
}

type person struct {
	name string
	age  int
}
type person4 struct {
	name string
	age  int
}

func main() {

	person1.name = "Test"
	person1.age = 25
	var person2 struct {
		name string
		age  int
	}

	person2.name = "Test"
	person2.age = 25

	person3 := struct {
		name string
		age  int
	}{
		name: "Test",
		age:  25,
	}
	fmt.Println(person2 == person3)
	fmt.Println(person1 == person2)
	fmt.Println(person1 == person3)

	person := person{name: "Test", age: 25}
	person4 := person4{name: "Test", age: 25}
	fmt.Println(person == person2)
	fmt.Println(person4)
	// fmt.Println(person == person4) // compilation error

}
