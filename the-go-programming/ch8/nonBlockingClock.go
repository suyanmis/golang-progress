package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("can't connect to the localhost:8000 over tcp")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("can't accept the connection to the localhost:8000 over tcp")
		}
		go handleConnection(conn)
	}
}

func handleConnection(con net.Conn) {
	defer con.Close()
	for {
		_, err := io.WriteString(con, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
