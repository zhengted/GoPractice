package main

import (
	"fmt"
)

// 责任链模式
// 以请假系统为例 day(请假天数) < 3 主管审核  3<day<=10 制作人审核  10<day<=30 总经理审核
// QAQ 对go的继承和虚函数实现还是不太熟悉

// 请求类
type leaveRequest struct {
	name string
	day int
}

// 抽象类
type Manager interface {
	checkDay(l leaveRequest) bool
	handleRequest(l leaveRequest) bool
}

type RequestChain struct {
	Manager
	nextHandler *RequestChain
}

func (r *RequestChain) SetNextHandler(m *RequestChain)  {
	r.nextHandler = m
}

func (r *RequestChain) handleRequest(l leaveRequest) bool  {
	if r.Manager.checkDay(l) {
		return r.Manager.handleRequest(l)
	}
	if r.nextHandler != nil {
		return r.nextHandler.handleRequest(l)
	}
	return false
}

func (r *RequestChain) checkDay(l leaveRequest) bool {
	return true
}

// 具体处理节点
// 直接主管
type DirectManager struct {}

func (r *DirectManager) checkDay(l leaveRequest) bool  {
	return l.day <= 3
}

func (r *DirectManager) handleRequest(l leaveRequest) bool  {
	fmt.Println("Direct manager permit")
	return true
}

func NewDirectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &DirectManager{},
	}
}



// 部门经理
type deptManager struct {}

func NewDeptManagerChain() *RequestChain {
	return &RequestChain{
		Manager : &deptManager{},
	}
}

func (r *deptManager) checkDay(l leaveRequest) bool  {
	return l.day > 3 && l.day <= 10
}

func (r *deptManager) handleRequest(l leaveRequest) bool  {
	fmt.Println("Dept manager permit")
	return true
}

// 总经理
type globalManager struct {}

func NewGlobalManagerChain() *RequestChain {
	return &RequestChain{
		Manager : &globalManager{},
	}
}

func (r *globalManager) checkDay(l leaveRequest) bool  {
	return l.day > 10 && l.day <= 30
}

func (r *globalManager) handleRequest(l leaveRequest) bool  {
	fmt.Println("Global manager permit")
	return true
}

func chainRepFactory() Manager {
	c1 := NewDirectManagerChain()
	c2 := NewDeptManagerChain()
	c3 := NewGlobalManagerChain()

	c1.SetNextHandler(c2)
	c2.SetNextHandler(c3)

	return c1
}

func TestChainOfResponsibility()  {
	c := chainRepFactory()
	l := leaveRequest{day: 40,name: "Bob"}
	c.handleRequest(l)
}

func main() {
	TestChainOfResponsibility()
}