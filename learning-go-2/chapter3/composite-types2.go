package main

import "fmt"

type slice []int

func main() {
	x := slice{1, 2, 3, 4}
	y := make([]int, 2)
	copy(y, x)
	fmt.Println(x, y)
	copy(y, x[2:])
	fmt.Println(x, y)
	// Array to slice
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1[:]
	arr2 = append(arr2, 5)
	fmt.Println(arr2)

	// array pointers
	arr3 := []int{1, 2, 3, 4}
	arr3Pointer := (*[4]int)(arr3)

	arr3Pointer[1] = 50
	arr3[0] = 15

	fmt.Println(arr3, "\n", arr3Pointer)

}
