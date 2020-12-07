package main

import (
	"fmt"
)

func CreateVector() Vector {
	s := []interface{}{}
	v := Vector{
		S: s,
	}
	return v
}

func vectorDemo() {
	v := CreateVector()
	v.Push_back(5)
	v.Push_back(3)
	v.Erase(0)
	fmt.Println(v)
}

func stackDemo() {
	st := Stack{
		S: []interface{}{},
	}
	fmt.Println(st.Empty())
	st.Push(1)
	st.Push(2)
	st.Push(3)
	fmt.Println(st.Empty())
	fmt.Println(st.Top())
	st.Pop()
	fmt.Println(st.Top())
	st.Pop()
	st.Pop()
	fmt.Println(st.Top())
}

func sideQueueDemo() {
	sq := SideQueue{
		Q: []interface{}{},
	}
	fmt.Println("side queue empty test")
	fmt.Println(sq.Empty())
	fmt.Println(sq.Back())
	fmt.Println(sq.Front())
	fmt.Println("side queue push test")
	sq.Push(1)
	sq.Push(2)
	sq.Push(3)
	fmt.Println(sq.Empty())
	fmt.Println(sq.Front())
	fmt.Println(sq.Back())
	fmt.Println("side queue pop test")
	sq.Pop()
	fmt.Println(sq.Empty())
	fmt.Println(sq.Front())
	fmt.Println(sq.Back())

}

func deQueueDemo() {

}

func main() {
	sideQueueDemo()
}
