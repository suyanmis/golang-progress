package main

import "fmt"

type Int interface {
	uint8 | uint16 | uint32 | uint64 | int16 | int32 | int64 | int | float32 | float64
}

type Printable[T int | float64] interface {
	String() T
}

func double[T Int](value T) T {
	return value * 2
}

type IntList[T int | float64] struct {
	list []T
}

func (sl IntList[T]) String() T {
	var total T
	for _, item := range sl.list {
		total += item
	}
	return total
}
func main() {
	fmt.Printf("double(5): %v\n", double(5))
	fmt.Printf("double(5.2): %v\n", double(5.2))
	il := IntList[int]{}
	il.list = []int{1, 3, 5, 6}
	fl := IntList[float64]{}
	fl.list = []float64{1.5, 2.5, 4.6, 6.4}

	fmt.Printf("il.String(): %v\n", il.String())
	fmt.Printf("fl.String(): %v\n", fl.String())
}
