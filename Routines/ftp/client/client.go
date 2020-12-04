package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Println(err.Error())
	}
	defer conn.Close()
	input := bufio.NewScanner(os.Stdin)
	go func() {
		for input.Scan() {
			_, err := io.WriteString(conn, string(input.Bytes()))
			if err != nil {
				log.Println(err.Error())
			}
		}
	}()

	//go func() {
	//	if _, err := io.Copy(os.Stdout, conn); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
}
