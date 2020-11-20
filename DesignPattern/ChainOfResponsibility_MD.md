## 概念
- 消息的接收者管理器为发送者创建了一系列的接收者链表
- 当接收者节点无法处理发送者的消息时将消息传递给下一个接收者
- **需要注意的点**：职责链中的每个节点需要实现同样的接口，保证职责链能够正常传递
- **优点**：请求的发送者和接收者解耦，发送者仅需知道职责链的头节点，将消息发送给头节点即可。链表的优点：允许动态增删，可以修改代码调整次序
- **缺点**：性能开销大，链表在查找过程是从前往后，不容易调试

- 代码：以下代码是以业务流中的请假作为示例
- 规则：
    - 请假天数小于等于3天，由直接主管审批
    - 请假天数小于等于10天，由项目主管审批
    - 请假天数小于等于30天，由总经理审批
    - 请假天数大于30天，不予审批
- 代码只是个demo还不太完善 见谅

```Golang
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
```