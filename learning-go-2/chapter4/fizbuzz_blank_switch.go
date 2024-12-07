package main

import (
	"fmt"
)

func main() {
	useGoTo()
}
func fizzBuzz(it int) {
	for i := 1; i <= it; i++ {
		switch { // blank switch
		case i%5 == 0 && i%3 == 0:
			fmt.Println("FizBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}

}
