package main

import (
	"fmt"
	"io"
)

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)
	fmt.Println(i.(MyInt) + 1)
	var i3 any
	i3 = "Test"
	fmt.Println(i3.(string) + " Test")
	// fmt.Println(i3 + "Test") fails
	i4, ok := i3.(int) // if ok is not fetched it will panic at runtime.
	if !ok {
		fmt.Printf("unexpected type for %v", i3)
	}
	fmt.Println(i4, ok)
}

// Type switch statements provide the ability to
// differentiate between multiple implementations of
// an interface that require different processing.
func doThings(i any) {
	switch j := i.(type) {
	case nil:

	case int:

	case MyInt:

	case io.Reader:

	case string:

	case bool, rune:

	default:
		fmt.Println(j)
	}
}
