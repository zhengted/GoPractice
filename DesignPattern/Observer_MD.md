### 观察者模式
- 行为型模式中重要程度Max的
- 被观察者执行了某类操作后，调用fire或者Notify 通知观察者修改状态
- 游戏中的实例非常多
    - 触发器（死亡触发器）
    - 成就系统
- 以触发器为例
    - 触发器列表会存放已经订阅该事件的怪物UID（怪物序列号也行）
    - 一旦死亡系统执行到死亡相关的事件时，向触发器列表发出消息
    - 触发器遍历执行（这里可以用迭代器模式）
- 代码
```golang
package main

import (
	"fmt"
)

/*
	观察者模式：行为型模式中最经常使用的一种
	意图：定义对象间的一种一对多的依赖关系，
		当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。
	游戏用途：触发器、成就系统
	成员：
		Subject：包含绑定观察者到Client对象和从Client对象解绑观察者的方法 （可以理解成玩家或系统行为）
		Observer：观察者类 （可以理解成成就系统Server）
*/
// 抽象观察者
type IObserver interface {
	Notify()
}

// 抽象被观察者
type ISubject interface {
	AddObservers(observers ...IObserver)	// 添加观察者
	NotifyObservers()						// 通知观察者
}

type Observer struct {}

func (o *Observer) Notify() {
	fmt.Println("已经触发了观察者")
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) AddObservers(observer ...IObserver) {
	s.observers = append(s.observers,observer...)
}

func (s *Subject) NotifyObservers() {
	for k := range s.observers {
		s.observers[k].Notify()
	}
}

func TestObserver() {
	s := new(Subject)
	o := new(Observer)

	s.AddObservers(o)
	o2 := new(Observer)
	s.AddObservers(o2)
	s.NotifyObservers()
}

func main() {
	TestObserver()
}
```