package main

import (
	"fmt"
	"sync"
)

func CreateWorker(id int,
	wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//done: make(chan bool),
		//wg : wg,
		done: func() {
			wg.Done()
		},
	}
	//go doWork(id, w.in, wg)
	go doWork(id, w)
	return w

}

func doWork(id int,
	w worker) {
	for n := range w.in {
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("Worker %d received %c number %d\n",
			id, n, n)
		//go func() {
		//	done <- true
		//}()
		w.done()
	}
}

type worker struct {
	in chan int
	// done	chan bool
	// wg		*sync.WaitGroup
	done func()
}

func chanDemo() {
	var (
		workers [10]worker
		wg      sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i // 阻塞式 第一个循环发了 正在等待done <- true
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
}

func main() {
	chanDemo()
}
