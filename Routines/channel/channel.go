package main

import (
	"fmt"
	"time"
)

func CreateWorker(id int) chan<- int { // <-chan：只发不收  chan<- ：只收不发
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {
	for n := range c {
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

func channelClose() {
	//  发送方close

}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	// chanDemo()
	fmt.Println("Buffered channnel")
	bufferedChannel() // 会因为close打印空串
	fmt.Println("Close channel")

	// 不要通过共享内存来通信；通过通信来共享内存
}
