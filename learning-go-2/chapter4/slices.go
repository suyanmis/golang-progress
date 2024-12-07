package main

import "fmt"

func main() {
	s := [...]int{1, 2, 3, 4, 5, 6}
	reverse(s[:])

	// fmt.Println(s)
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotate(array)
	fmt.Println(array)

	arraystr := []string{"a", "b", "c", "d", "e", "a", "b", "c", "d", "e", "a", "b", "q"}
	deleteDuplicates(&arraystr)
	fmt.Println(arraystr)

}

func reverse(array []int) {
	for i, j := 0, len(array)-1; j > i; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func rotate(array []int) {
	// remainder := make([]int, len(array))
	intVal := array[len(array)-1]
	// copy(remainder, array)
	// array[0] = array[len(array)-1]
	copy(array[1:], array)
	array[0] = intVal
}
func deleteDuplicates(array *[]string) {
	set := make(map[string]bool)
	newArray := make([]string, 0, len(*array))
	for _, s := range *array {
		if !set[s] {
			set[s] = true
			newArray = append(newArray, s)
		}
	}
	*array = newArray
}
