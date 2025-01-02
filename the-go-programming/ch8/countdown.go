package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick:
			// DN
			fmt.Printf("\r%2d", countdown)
		case <-abort:
			aborting()
		}
	}
	launch()
}

func launch() {
	fmt.Println("\rLaunching!!!")
}
func aborting() {
	log.Fatalln("\rAborting!")
}
