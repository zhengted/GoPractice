package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier // ducktyping 使用者（simple queue）实现了就行
	Submit(Request)
	WorkerChan() chan Request // 询问给哪个worker
	Run()
}

type ReadyNotifier interface {
	WorkReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// simple调度的写法
	//in := make(chan Request)
	//out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	//fmt.Println("Init scheduler")
	//
	//for i := 0; i < e.WorkerCount; i++ {
	//	createWorker(in,out)
	//}
	//
	//for _,r := range seeds {
	//	e.Scheduler.Submit(r)
	//}

	// 队列调度的写法
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if IsDuplicate(r.Url) {
			log.Printf("Duplicate request:"+"%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		// URL dedup
		for _, request := range result.Requests {
			if IsDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request) // 这里会有个问题  如果result里的request数量过大会循环等待 解决办法是开一个routine
		}
	}
}

// Simple调度的写法
//func createWorker(in chan Request, out chan ParseResult) {
//	go func() {
//		for  {
//			// tell scheduler i'm ready
//			request := <- in	// *1
//			result, err := worker(request)
//			if err != nil {
//				continue
//			}
//			out <- result		// *1和这里循环等待了
//		}
//	}()
//}

// 队列调度的写法
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, notifier ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			notifier.WorkReady(in)
			request := <-in // *1
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result // *1和这里循环等待了 解决办法是submit起一个goroutine
		}
	}()
}

// 去重
var visitedUrls = make(map[string]bool)

func IsDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
