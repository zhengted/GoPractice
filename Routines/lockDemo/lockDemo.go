package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		for j := 0; j < 100; j++ {
			total.value += 1
		}
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go worker(&wg)
	go worker(&wg)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.value)
}
