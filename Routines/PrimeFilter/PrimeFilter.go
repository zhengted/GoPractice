package main

import (
	"context"
	"fmt"
)

// 素数筛实现 普通版

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			fmt.Printf("Insert value to channel: %v\n", i)
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				fmt.Printf("Get next prime : %v\n", i)
				out <- i
			}
		}
	}()
	return out
}

func main() {
	//ch := GenerateNatural()
	//for i := 0; i < 100; i++ {
	//	prime := <-ch // 每从管道中读取一个数就将他作为筛子 去寻找下一个素数
	//	fmt.Printf("%v: %v\n", i+1, prime)
	//	ch = PrimeFilter(ch, prime) // 这里获取到的新ch是被筛选过的数据
	//}
	mainEx()
}

//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch1-basic/ch1-06-goroutine.md

// 素数筛实现增强版
// 引入context包  针对线程安全退出和超时控制
// 以素数筛简单版为例，当main不使用两个go routine之后会存在内存泄漏的风险
func GeneratorNaturalEx(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func PrimeFilterEx(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func mainEx() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := GeneratorNaturalEx(ctx)
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilterEx(ctx, ch, prime)
	}
	cancel()
	// 通过context来维持后台Goroutine的状态
}
