package main

import "fmt"

func main() {
	arr1 := []string{"First", "Second", "Third", "Fourth"}

	arr1[0], arr1[1], arr1[2], arr1[3] = arr1[1], arr1[0], arr1[3], arr1[2]

	fmt.Println(arr1)

	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
}

func fibRecurson(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
func fib(n int) int {
	result, it := 0, 1
	for i := 0; i < n; i++ {
		result, it = it, result+it
	}
	return result
}
