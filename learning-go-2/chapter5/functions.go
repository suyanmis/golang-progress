package main

import (
	"errors"
	"fmt"
)

var (
	add = func(i, j int) int { return i + j }
	sub = func(i, j int) int { return i * j }
	mul = func(i, j int) int { return i / j }
	div = func(i, j int) int { return i - j }
)
var i1 int = 5

func main() {
	fmt.Println(addTo(3, 5, 1, 2, 3, 4, 5, 6))

	fmt.Println(divAndRemainder(10, 3))

	divAndRem(10, 3)
	// Closures
	sideEffect()
	fmt.Println(i1)

	var dAndR func(i1, i2 int)
	dAndR = divAndRem

	dAndR(10, 3)

	add := func(a1 int, a2 int) int {
		return a1 + a2
	}

	fmt.Println(add(3, 5))

	fmt.Println(func(i1 int, i2 int) int {
		return i1 + i2
	}(3, 5))

	mul(3, 5)
	div(5, 3)
	sub(5, 10)
}

func addTo(init int, integers ...int) []int {
	slice := make([]int, 0, len(integers))
	for _, value := range integers {
		slice = append(slice, value+init)
	}
	return slice
}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

var divAndRem = func(num, denom int) {
	if denom == 0 {
		fmt.Println("cannot divide by zero")
		return
	}
	fmt.Println(num/denom, num%denom)
}

type opFuncType func(int, int) int

var opMap map[string]opFuncType

func sideEffect() {
	fmt.Println(i1)
	i1 = 10

}
