package main

import "fmt"

func main() {
	nat := make(chan int)
	squ := make(chan int)

	go counter(10, nat)
	go squarer(nat, squ)
	list := make([]int, 0, 10)
	// Option1
	// for {
	// 	item, ok := <-squ
	// 	if !ok {
	// 		break
	// 	}
	// 	list = append(list, item)
	// }

	// Option 2
	for i := range squ {
		list = append(list, i)
	}
	fmt.Println(list)

}
func squarer(in <-chan int, out chan<- int) {
	// Option 1
	// for {
	// 	x, ok := <-nat
	// 	if !ok {
	// 		break
	// 	}
	// 	squ <- x * x
	// }
	// close(squ)

	// Option 2
	for i := range in {
		out <- i * i
	}
	close(out)
	fmt.Println("Finished squarer")
}
func counter(i int, out chan<- int) {
	for count := 0; count < i; count++ {
		out <- count
	}
	close(out)
	fmt.Println("Finished counter")
}
