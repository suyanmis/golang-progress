package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(fn string, ln string, age int) Person {
	return Person{FirstName: fn, LastName: ln, Age: age}
}

func MakePersonPointer(fn string, ln string, age int) *Person {
	return &Person{FirstName: fn, LastName: ln, Age: age}
}
func main() {
	// MakePerson("Serhat", "uyanmis", 28)
	// MakePersonPointer("Serhat", "uyanmis", 28)
	s := []string{"test", "me", "first", "or-not"}
	v := "last"
	UpdateSlices(s, v)
	fmt.Println(s)
	GrowSlice(s, "new")
	fmt.Println(s)

}

func UpdateSlices(s []string, v string) {
	s[len(s)-1] = v
}

func GrowSlice(s []string, v string) {
	s = append(s, v) // without return it doesn't work
}
