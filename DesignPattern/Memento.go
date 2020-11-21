package main

import "fmt"

/*
	备忘录模式
	在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。
	举例：塞尔达自动存档，Ctrl+z操作
	三个类：
		Originator	存储状态的类 创建并在Memento对象中存储状态
		Memento		状态类 包含了要被回复的对象的状态
		CareTaker	负责恢复状态的类
*/

type Memento struct {
	state string
}

func (m *Memento) Memento(state string) *Memento {
	m.state = state
	return m
}

func (m *Memento) getState() string {
	return m.state
}

type Originator struct {
	state string
}

func (o *Originator) setState(state string) {
	o.state = state
}

func (o *Originator) getState() string {
	return o.state
}

func (o *Originator) saveStateToMemento() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) getStateFromMemento(memento Memento)  {
	o.state = memento.state
}

type CareTaker struct {
	mementoList []*Memento

}

func (t *CareTaker) add(memento *Memento)  {
	t.mementoList = append(t.mementoList, memento)
}

func (t *CareTaker) getBack() *Memento {
	ret := t.mementoList[len(t.mementoList) - 1]
	t.mementoList = t.mementoList[:len(t.mementoList)-1]
	return ret
}

func TestMemento()  {
	o := Originator{state: "1"}
	c := CareTaker{}
	o.setState("2")
	c.add(o.saveStateToMemento())	// 存档
	fmt.Println("存档操作2 当前状态",o.getState())
	o.setState("3")
	c.add(o.saveStateToMemento())	// 存档
	fmt.Println("存档操作3 当前状态",o.getState())

	o.setState("4")
	o.setState("5")
	fmt.Println("无存档操作5 当前状态",o.getState())
	// 回档操作
	m := c.getBack()
	o.setState(m.getState())
	fmt.Println("回档操作 当前状态",o.getState())

}
func main()  {
	TestMemento()
}