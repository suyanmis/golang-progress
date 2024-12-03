package main

import (
	"ch2/temperature"
	"fmt"
)

// package-level declarations
const boilingC = 100.0

func main() {
	// local declarations
	var c = boilingC
	var f = (c * 9 / 5) + 32
	fmt.Println(temperature.CtoF(temperature.Celcius(c)))
	fmt.Printf("Boiling point= %g F or %g C\n", f, c)

}
