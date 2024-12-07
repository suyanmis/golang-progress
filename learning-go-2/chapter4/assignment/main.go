package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intSlice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		intSlice = append(intSlice, rand.Intn(100))
	}
	for _, i := range intSlice {
		switch {
		case i%3 == 0 && i%2 == 0:
			fmt.Println("SIX")
		case i%2 == 0:
			fmt.Println("TWO")
		case i%3 == 0:
			fmt.Println("THREE")
		default:
			fmt.Println("Never mind")
		}
	}
	var total int

	for i := 0; i < 10; i++ {
		total := total + i // side effect due to the scope, reassigns the total value.
		// total = total + i
		fmt.Println(total)
	}

}
