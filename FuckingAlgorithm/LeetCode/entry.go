package main

import (
	"fmt"
	"log"
	"strconv"
)

var t int

func init() {
	t, err := strconv.Atoi("2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("init:", t)
}

func main() {
	fmt.Println("main:", t)
}
