package main

import "fmt"

/*
	空对象模式
		利用一个空对象取代Null对象实例的检查
*/

type AbstractCustomer interface {
	isNil() bool
	getName() string
}

type RealCustomer struct {
	name string
}

func (r *RealCustomer) RealCustomer(name string) *RealCustomer {
	r.name = name
	return r
}

func (r *RealCustomer) isNil() bool {
	return false
}

func (r *RealCustomer) getName() string {
	return r.name
}

type NullCustomer struct {

}

func (n *NullCustomer) isNil() bool {
	return true
}

func (n *NullCustomer) getName() string {
	return "Not Available in Customer Database"
}

// 重点 工厂类怎么运作
type CustomerFactory struct {
	names []string
}

func (cf *CustomerFactory) add(names ...string)  {
	cf.names = append(
		cf.names,
		names...,
	)
}

func (cf *CustomerFactory) getCustomer(name string) AbstractCustomer {
	for _,n := range cf.names {
		if(n == name) {
			return &RealCustomer{name: n}
		}
	}
	return &NullCustomer{}
}

func TestNullObject()  {
	cf := CustomerFactory{}
	cf.add("bob","john","King","Nine")

	c1 := cf.getCustomer("bob")
	fmt.Printf("Is c1 Nil ?\n",c1.isNil())
	fmt.Printf("c1 name %s\n",c1.getName())

	c2 := cf.getCustomer("Adele")
	fmt.Printf("Is c2 Nil ?\n",c2.isNil())
	fmt.Printf("c2 name %s\n",c2.getName())
}

func main()  {
	TestNullObject()
}