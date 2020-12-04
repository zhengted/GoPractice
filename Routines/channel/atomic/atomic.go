package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
	sync.Map
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func() { // 实现代码区加锁
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()

}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
