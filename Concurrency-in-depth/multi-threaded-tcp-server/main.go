package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	conn.Write([]byte("HTTP/1.0 200 OK\r\n\r\n Hello, World!\r\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("Waiting for a client to connect")
		connection, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Processing the request")

		go do(connection)
	}
}
