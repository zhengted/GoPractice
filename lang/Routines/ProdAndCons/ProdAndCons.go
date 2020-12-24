package main

import "fmt"

func main() {
	data := make(chan int)
	done := make(chan bool)
	go Consumer(data, done)
	go Producer(data)
	<-done
}

func Consumer(data chan int, done chan<- bool) {
	for i := range data {
		fmt.Printf("Consumer get: %d\n", i)
	}
	done <- true
}

func Producer(data chan int) {
	defer close(data)
	for i := 0; i < 4; i++ {
		fmt.Printf("Producer produce num : %d\n", i)
		data <- i
	}
}
