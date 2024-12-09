package main

import (
	"fmt"
	"sort"
)

func main() {
	var ages map[string]int

	ages["carol"] = 21

	if res, ok := ages["carol"]; !ok {
		fmt.Println("Key doesn't exist!")
	} else {
		fmt.Println(res)
	}
	delete(ages, "test") // no-op
	// map initialization

	for name := range ages {
		fmt.Println("Name:", name)
	}
	for name, age := range ages {
		fmt.Println("Name:", name, "Age:", age)
	}
	// hashmaps are always unordered
	// to sort the names
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

}
