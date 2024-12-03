package main

import "fmt"

type set[T comparable] interface {
	union(s set[T]) set[T]
	intersect(s set[T]) set[T]
}
type stringSet map[string]bool

func (s stringSet) union(set stringSet) stringSet {
	result := make(stringSet)
	for k, _ := range set {
		result[k] = true
	}
	for k, _ := range s {
		result[k] = true
	}
	return result
}

func (s stringSet) intersect(set stringSet) stringSet {
	result := make(stringSet)
	if len(s) > len(set) {
		for k, _ := range set {
			if _, ok := s[k]; ok {
				result[k] = true
			}
		}
	} else {
		for k, _ := range s {
			if _, ok := set[k]; ok {
				result[k] = true
			}
		}
	}

	return result
}
func main() {
	set1 := make(stringSet)
	set2 := make(stringSet)
	arr1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "a", "b", "c", "d", "e", "f", "g", "h"}
	for _, value := range arr1 {
		set1[value] = true
	}
	arr2 := []string{"a", "b", "c", "1", "2", "3", "4", "5", "6", "e"}
	for _, value := range arr2 {
		set2[value] = true
	}
	union := set1.union(set2)
	intersect := set1.intersect(set2)

	fmt.Println(union)
	fmt.Println(intersect)

}

func union[T comparable](set1, set2 set[T]) set[T] {
	return set1.union(set2)
}

func intersect[T comparable](set1, set2 set[T]) set[T] {
	return set1.intersect(set2)
}
