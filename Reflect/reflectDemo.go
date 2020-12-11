package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId    int
	customer int
}

type employee struct {
	id   int
	name string
	Sex  string
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	fmt.Println("Type:", t)
	fmt.Println("Value:", v)
	fmt.Println("Kind:", k)
	fmt.Printf("%T\n", q)
	fmt.Println(q)
}

func main() {
	o := order{
		456, 56,
	}
	createQuery(o)
	e := employee{
		1,
		"killer",
		"male",
	}
	createQuery(e)
}
