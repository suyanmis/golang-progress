package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go io.Copy(os.Stdout, conn)
	io.Copy(conn, os.Stdin)
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "\t", strings.Title(strings.ToLower(shout)))
	time.Sleep(delay)
	fmt.Fprintf(c, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}
func handleConnection(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 500*time.Millisecond)
	}
	if input.Err() == nil || input.Err() == io.EOF {

	} else {
		log.Fatal(input.Err().Error())
	}
	c.Close()
}
