package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "loaclhost:8888")
	if err != nil {
		log.Println(err.Error())
	}

	defer conn.Close()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if _, err := conn.Write(input.Bytes()); err != nil {
			log.Println(err.Error())
		}
	}

}

func copyToReader(src io.Writer, dst io.Reader) {
	if _, err := io.Copy(src, dst); err != nil {
		log.Println(err.Error())
	}
}
