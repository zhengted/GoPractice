package main

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
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
