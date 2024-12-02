package main

import "fmt"

// package-level declarations
const boilingC = 100.0

func main() {
	// local declarations
	var c = boilingC
	var f = (c * 9 / 5) + 32

	fmt.Printf("Boiling point= %g F or %g C\n", f, c)

}
