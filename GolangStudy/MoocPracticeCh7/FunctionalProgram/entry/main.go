package main

import (
	"fmt"
)

/*
	函数式编程VS函数指针
	- 函数是一等公民：参数，变量，返回值都可以是函数
	- 高阶函数

	“正统”的函数编程
	- 不可变性：不能有状态，只有常量和函数
	- 函数只能有一个参数

*/

func adder() func(int) int {
	sum := 0	// upvalue 自由变量环境
	// 返回的不仅仅是函数 还有自由变量
	return func(v int) int {
		sum += v	// v 局部变量
		return sum
	}
}

//"正统"的函数式编程
type iAdder func(int) (int, iAdder)
func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base+v,adder2(base+v)
	}
}

func main() {
	a := adder2(0)

	for j := 0; j < 10; j++ {
		var s int
		s, a = a(j)
		fmt.Println(j,s)
	}

}
