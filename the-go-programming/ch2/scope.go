package main

import (
	"fmt"
	"os"
)

func main() {
	x := "hello"
	fmt.Print('A' == 'a')
	fmt.Print('t' - 't')

	for _, x := range x {
		x := x + 'a' - 'A'
		fmt.Printf("%c", x)
	}

	if x := f(); x == 0 {
		fmt.Println("Test")
	}

	if f, err := os.Open("test.txt"); err == nil {
		fmt.Printf("f: %v\n", f)
	}
}
