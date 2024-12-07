package main

import (
	"fmt"
	"math/rand"
)

func main() {

}

// This is proper way of using goto. It breaks the infinite loop.
func useGoTo() {
	a := rand.Intn(10)
	for a < 100 {
		fmt.Println("Number: ", a)
		if a%5 == 0 {
			goto final
		}
		a = rand.Intn(10)
	}
final:
	fmt.Println("Triggered ", a)
}
