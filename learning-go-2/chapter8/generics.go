package main

import (
	"fmt"
	"log"
)

type Stack[T comparable] struct {
	vals []T
}

// if it's any type it won't work
// type Stack[T any] struct {
// 	vals []T
// }

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val) // append value to the struct
}

func (s Stack[T]) Contains(val T) bool {
	for _, item := range s.vals {
		if val == item {
			return true
		}
	}
	return false
}

// LIFO
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]    // Last item
	s.vals = s.vals[:len(s.vals)-1] // Last item -1
	return top, true
}
func main() {
	stack := Stack[int]{vals: []int{1, 2, 3, 4, 5, 6}}
	fmt.Printf("stack.Contains(5): %v\n", stack.Contains(5))
	for len(stack.vals) > 0 {
		_, ok := stack.Pop()
		// Never triggers here.
		if !ok {
			log.Println("Empty")
			break
		}
		// log.Printf("Item: %d\n", item)
	}
	stack.Push(10)
	stack.Push(20)
	i, _ := stack.Pop()
	fmt.Println(i)
	i, _ = stack.Pop()
	fmt.Println(i)
	i, ok := stack.Pop()
	if !ok {
		fmt.Println("Empty array. Returning default:", i)
	}
	log.Printf("%#v", stack)

	filteredList := Filter([]string{"adam", "barby", "cam", "ataman", "abi"}, func(val string) bool {
		if val[0] == 'a' {
			return true
		}
		return false
	})
	fmt.Println("Filtered list:", filteredList)
	reducedList := Reduce(filteredList, "test", func(acc string, val string) string {
		return acc + val + ","
	})
	reducedList = reducedList[:len(reducedList)-1]
	fmt.Println("Reduced list:", reducedList)

}

// Filter any value within the given array for the given function
func Filter[T any](items []T, f func(val T) bool) []T {
	var list []T
	for _, item := range items {
		if f(item) {
			list = append(list, item)
		}
	}
	return list
}

// Reduces the given array to a single object
func Reduce[T1, T2 any](items []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, item := range items {
		r = f(r, item)
	}
	return r
}
