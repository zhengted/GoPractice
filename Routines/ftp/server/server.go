package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Println(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println("Connect Success:", conn.LocalAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		var temp []byte
		if _, err := conn.Read(temp); err != nil {
			return
		}
		if len(temp) <= 0 {
			continue
		}
		fmt.Println(temp)
		time.Sleep(time.Second)
	}
}
