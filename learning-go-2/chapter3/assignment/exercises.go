package main

import "fmt"

type Employee struct {
	id        int64
	firstName string
	lastName  string
}

func main() {
	greetWords := []string{"Hello", "Ola", "こにちは", "Привіт", "नमस्कार"}
	firstTwo := greetWords[0:2]
	secondFourth := greetWords[2:5] // possible [2:]
	fourthFifth := greetWords[3:5]  // possible [3:5]

	fmt.Println(firstTwo, secondFourth, fourthFifth)

	var employee1 Employee
	employee1.firstName = "John"
	employee1.lastName = "Doe"
	employee2 := Employee{}
	employee3 := Employee{firstName: "John", lastName: "Doe"}

	fmt.Println(employee2, employee3)
}
