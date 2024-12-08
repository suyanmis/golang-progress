package main

import (
	"fmt"
	"sort"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := Person{"Serhat", "Uyanmis", 23}
	p2 := Person{"Abuzer", "Kadayif", 26}
	p3 := Person{"Tahammul", "Edemiyorum", 28}

	people := make([]Person, 0, 3)

	people = append(people, p1, p2, p3)

	sort.Slice(people, func(i, j int) bool {
		return people[i].firstName <= people[j].firstName
	})
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println(people)
}
