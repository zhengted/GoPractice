package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func version1() {
	var mu sync.Mutex
	go func() {
		fmt.Println("Hello,World!")
		mu.Lock() // 和Unlock并发执行，如果Unlock先执行会发生异常
	}()
	mu.Unlock()
}

//因为mu.Lock()和mu.Unlock()并不在同一个Goroutine中，
//所以也就不满足顺序一致性内存模型。同时它们也没有其它的同步事件可以参考，
//这两个事件不可排序也就是可以并发的。因为可能是并发的事件，
//所以main函数中的mu.Unlock()很有可能先发生，
//而这个时刻mu互斥对象还处于未加锁的状态，从而会导致运行时异常。

func version2() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("Hello,world")
		mu.Unlock()
	}()
	// 执行到这个Lock会被阻塞，直到线程中的mu被Unlock
	mu.Lock()
}

//修复的方式是在main函数所在线程中执行两次mu.Lock()，
//当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，
//main函数的阻塞状态驱动后台线程继续向前执行。
//当后台线程执行到mu.Unlock()时解锁，此时打印工作已经完成了，
//解锁会导致main函数中的第二个mu.Lock()阻塞状态取消，
//此时后台线程和主线程再没有其它的同步事件参考，
//它们退出的事件将是并发的：在main函数退出导致程序退出时，后台线程可能已经退出了，也可能没有退出。
//虽然无法确定两个线程退出的时间，但是打印工作是可以正确完成的。

func version3() {
	done := make(chan int)
	go func() {
		fmt.Println("Hello,world!")
		<-done
	}()
	done <- 1
}

//根据Go语言内存模型规范，对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。
//因此，后台线程<-done接收操作完成之后，main线程的done <- 1发送操作才可能完成（从而退出main、退出程序），
//而此时打印工作已经完成了。
//上面的代码虽然可以正确同步，但是对管道的缓存大小太敏感：
//如果管道有缓存的话，就无法保证main退出之前后台线程能正常打印了。

func version4() {
	done := make(chan int, 1)
	go func() {
		fmt.Println("Hello,world!")
		done <- 1
	}()
	<-done
}

//对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前，
//其中C是Channel的缓存大小。
//虽然管道是带缓存的，main线程接收完成是在后台线程发送开始但还未完成的时刻，
//此时打印工作也是已经完成的。

func version5() {
	done := make(chan int, 10)
	for i := 0; i < cap(done); i++ {
		//index := i
		go func(index int) {
			fmt.Println("Hello,World", index)
			//fmt.Println("Hello,World",i)
			done <- 1
		}(i)
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}
}

// 基于带缓存的管道，我们可以很容易将打印线程扩展到N个。
// 注意循环里起若干个协程写法要做出改变参考注释

func version6() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Hello,World!")
			wg.Done()
		}()
	}
	wg.Wait()
}

//其中wg.Add(1)用于增加等待事件的个数，必须确保在后台线程启动之前执行（如果放到后台线程之中执行则不能保证被正常执行到）.
//当后台线程完成打印工作之后，调用wg.Done()表示完成一个事件。
//main函数的wg.Wait()是等待全部的事件完成。

// 生产者消费者并发
// 生产者
func Producer(factor int, out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func ProdConsuDemo() {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)

	time.Sleep(5 * time.Second)
}

func ProdConsuDemoEx() {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)

	// Ctrl+c 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

func main() {
	fmt.Println("Test cocurrency")
	ProdConsuDemoEx()
}
