package main

import "fmt"

type IVector interface {
	Push_back(v interface{})
	Pop_back()
	Front() interface{}
	Back() interface{}
	Empty() bool
	Size() int
	Insert(v interface{}, index int)
	Erase(index int)
}

type Vector struct {
	s []interface{}
}

func CreateVector() Vector {
	s := []interface{}{}
	v := Vector{
		s: s,
	}
	return v
}

func (this *Vector) Push_back(v interface{}) {

	this.s = append(this.s, v)
}

func (this *Vector) Pop_back() {

	this.s = this.s[:len(this.s)-1]

}

// e.g. temp := v.Front
func (this *Vector) Front() interface{} {
	if len(this.s) <= 0 {
		return nil
	}
	return this.s[0]
}

func (this *Vector) Back() interface{} {
	if len(this.s) <= 0 {
		return nil
	}
	return this.s[len(this.s)-1]
}

func (this *Vector) Empty() bool {
	return len(this.s) <= 0
}

func (this *Vector) Size() int {
	return len(this.s)
}

func (this *Vector) Insert(value interface{}, nIndex int) {
	temp := this.s[nIndex : len(this.s)-1]
	this.s = this.s[:nIndex]
	this.s = append(this.s, value)
	this.s = append(temp)
}

func (this *Vector) Erase(nIndex int) {
	this.s = append(this.s[:nIndex], this.s[nIndex+1:])
}

func main() {
	v := CreateVector()
	v.Push_back(5)
	v.Push_back(3)
	v.Erase(0)
	fmt.Println(v)
}
