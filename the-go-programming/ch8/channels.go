package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// connection between go routines
	// equal only when two channel ref same data structure.
	// send data 	|   <- data
	// receive data |	<-ch
	// close(ch)
	// if no value initialized unbuffered chan
	ch := make(chan int)             // zero value == nil
	chBuffered := make(chan int, 10) // capacity of 10

	//

}

func dialHttp() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// not necessarily needs to be struct can be done by bool or int too.
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} //<--------------------|
	}()
	io.Copy(conn, os.Stdin)
	conn.Close()
	<-done // wait for go routine to finish first.|
}
