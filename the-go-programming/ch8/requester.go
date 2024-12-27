package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("no connection")
	}
	defer conn.Close()

	io.Copy(os.Stdout, conn)
}
