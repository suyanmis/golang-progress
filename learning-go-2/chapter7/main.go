package main

import "fmt"

type Person struct {
	FirstName string
	Age       int
}

// receiver as pointers
func (p *Person) changeAge(age int) {
	p.Age = age
}

func main() {
	p := Person{FirstName: "Test"}
	p.changeAge(35)

	fmt.Println(p)

	m := Manager{}
	m.Name = "Bruce" // direct access to the fields
	m.ID = "1234"
	m.Employee.Name = "Brian"

	fmt.Println(m.Name, m.Employee.Name)
}

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee //embedded fields
	reports  []Employee
	Name     string
}
