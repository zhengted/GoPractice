package main

type IQueue interface {
	Empty() bool
	Size() int
	Front() interface{}
	Back() interface{}
}

// 单向队列
type SideQueue struct {
	Q []interface{}
}

func (sq *SideQueue) Empty() bool {
	return len(sq.Q) <= 0
}

func (sq *SideQueue) Size() int {
	return len(sq.Q)
}

func (sq *SideQueue) Front() interface{} {
	if sq.Empty() {
		return nil
	}
	return sq.Q[0]
}

func (sq *SideQueue) Back() interface{} {
	if sq.Empty() {
		return nil
	}
	return sq.Q[len(sq.Q)-1]
}

func (sq *SideQueue) Push(val interface{}) {
	sq.Q = append(sq.Q, val)
}

func (sq *SideQueue) Pop() {
	sq.Q = sq.Q[1:]
}

// 双向队列
type DeQueue struct {
	Q []interface{}
}

func (dq *DeQueue) Empty() bool {
	return len(dq.Q) <= 0
}

func (dq *DeQueue) Size() int {
	return len(dq.Q)
}

func (dq *DeQueue) Front() interface{} {
	return dq.Q[0]
}

func (dq *DeQueue) Back() interface{} {
	return dq.Q[len(dq.Q)-1]
}

func (dq *DeQueue) Push_back(val interface{}) {
	dq.Q = append(dq.Q, val)
}

func (dq *DeQueue) Push_front(val interface{}) {
	rear := append([]interface{}{}, dq.Q[:len(dq.Q)])
	dq.Q = append([]interface{}{}, val)
	dq.Q = append(dq.Q, rear...)
}

func (dq *DeQueue) Pop_back() {
	dq.Q = dq.Q[:len(dq.Q)-1]
}

func (dq *DeQueue) Pop_front() {
	dq.Q = dq.Q[1:]
}
