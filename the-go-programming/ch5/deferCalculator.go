package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// runSleep()
	// Execute before panic to activate recovery.

	panicing() // safely recovered by using defer func(  recover() )()

}

func runSleep() {
	defer trace("runSleep")() // entry and end
	time.Sleep(2 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("Entering %s", msg)
	return func() { log.Printf("exit %s(%s)", msg, time.Since(start)) }
}

func panicing() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering the panic:", r)
		}
	}()
	fmt.Println("Ready to panic!!!!")
	panic("OMG!")
	fmt.Println("It will not work.")
}
