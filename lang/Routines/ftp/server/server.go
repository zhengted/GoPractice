package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Println(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		fmt.Println("Connect Success:", conn.LocalAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			return // e.g., client disconnect
		}

	}
}
