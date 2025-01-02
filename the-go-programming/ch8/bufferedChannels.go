package main

import "fmt"

func main() {
	ch := make(chan string, 3) // buffered channel

}

func m(c chan<- int) {
	c <- 3
}
func mn(c <-chan int) {
	v := <-c

	fmt.Println(v)
}
