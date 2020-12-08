package Scheduler

import "GoPractice/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request //	Worker的类型是 chan engineRequest
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// 告知Q 有一个channel（Worker）已经准备好工作了
func (q *QueuedScheduler) WorkReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueuedScheduler) Run() {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				// 这里不执行发 会导致下面的r和w收不到东西
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			// 这两个先后顺序不固定 用select控制
			case r := <-q.requestChan:
				// send r to a unknown worker
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				// send unknown next request to w
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}
	}()
}
