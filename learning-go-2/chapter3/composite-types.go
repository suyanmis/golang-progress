package main

import "fmt"

// If possible don't increase the capacity of the slices(hard copy)
// arrays are simple, it is fixed size.
func arrayInit() {
	intArray := make([]int, 0, 10)

	fmt.Println(intArray)
	appendInt(&intArray, 10, 20, 30, 40)
	fmt.Println(intArray)
	clear(intArray)
	fmt.Println(intArray)

	var nilSlice []string
	var emptySlice = []string{}
	var emptyArray = [10]string{}
	fmt.Printf("nilSlice %v has length %d and capacity %d\n", nilSlice, len(nilSlice), cap(nilSlice))
	fmt.Printf("emptySlice %v has length %d and capacity %d\n", emptySlice, len(emptySlice), cap(emptySlice))
	fmt.Printf("emptyArray %v has length %d and capacity %d\n", emptyArray, len(emptyArray), cap(emptyArray))

	config := []string{"go", "run", "."}
	// safest approach
	safeSlice := make([]string, 0)
	// assigning 0 length will make it behave like a nil object but when it is required to extend the slice,
	// it will dynamically allocate a new slice with a bigger capacity and non zero elements.
	fmt.Println(config)
	fmt.Println(safeSlice)
	fmt.Printf("Zero length slice is equal to null: %v\n", safeSlice == nil) //asigned
	fmt.Printf("Empty slice is equal to null: %v\n", emptySlice == nil)      //asigned
	fmt.Printf("nil slice is equal to null: %v\n", nilSlice == nil)          // no assignment

	configReference := config[0:2]

	fmt.Println(configReference)

	configReference[0] = "Changed value"

	fmt.Printf("config[0] equal to 'Changed value': %v", config[0] == "Changed value")
	//subslices pass by reference
}

func appendInt(array *[]int, values ...int) {
	*array = append(*array, values...)
}
