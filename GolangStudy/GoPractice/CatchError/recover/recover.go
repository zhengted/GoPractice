package main

import (
	"fmt"
)

func revoverTest()  {
	defer func() {
		e := recover()
		if err,ok := e.(error); ok {
			fmt.Println("error occured",err)
		}else {
			panic(err)
		}

	}()
	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {
	revoverTest()
}
