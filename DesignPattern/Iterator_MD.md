### 迭代器模式
- 就是个教你写迭代器的模式
- 建议去看STL iterator的实现文档
- 我这代码没啥意义
- 代码
```golang
package main

import (
	"container/list"
	"fmt"
)

/*
	迭代器模式
	意图：遍历对象
*/
type Container interface {
	Iterator() Iterator
}

type List struct {
	list list.List
	length int
}

func (l *List) Iterator() Iterator {
	return &ListIterator{l.list.Front(),l.list.Back()}
}

func (l *List) Add(value interface{})  {
	l.list.PushBack(value)
	l.length += 1
}

type Iterator interface {
	HasNext() bool
	Value() interface{}
	Next()
}

type ListIterator struct {
	cur *list.Element
	end *list.Element
}

func (li *ListIterator) HasNext() bool {
	return li.cur != li.end
}

func (li *ListIterator) Value() interface{} {
	return li.cur.Value
}

func (li *ListIterator) Next()  {
	li.cur = li.cur.Next()
}

func TestIteratorPattern()  {
	l := List{}
	l.Add("aaa")
	l.Add("bbb")
	l.Add("ccc")
	l.Add("ddd")
	i := l.Iterator()
	for i.HasNext() {
		x := i.Value()
		fmt.Println(x)
		i.Next()
	}
}

func main() {
	TestIteratorPattern()
}
```