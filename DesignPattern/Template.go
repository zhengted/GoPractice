package main

import "fmt"

/*
	模板模式
		定义一个操作中的算法的骨架，而将一些步骤延迟到子类中。
		模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
	举例
		1.spring 中对 Hibernate 的支持，将一些已经定好的方法封装起来，比如开启事务、获取 Session、关闭 Session 等，
	程序员不重复写那些已经规范好的代码，直接丢一个实体就可以保存。
		2.副本流程  副本初始虎啊、副本开始、副本结束
*/

type Game struct {
	init func()
	startPlay func()
	endPlay func()
}

func (g *Game)Play() {
	g.init()
	g.startPlay()
	g.endPlay()
}

type Cricket struct {
	Game
}

func (c *Cricket) init()  {
	fmt.Println("Cricket init")
}

func (c *Cricket) startPlay()  {
	fmt.Println("Cricket start Play ")
}

func (c *Cricket) endPlay()  {
	fmt.Println("Cricket end Play")
}

func NewCricket() *Cricket {
	cr := new(Cricket)
	cr.Game.init = cr.init
	cr.Game.startPlay = cr.startPlay
	cr.Game.endPlay = cr.endPlay
	return cr
}

type Football struct {
	Game
}

func (f *Football) init()  {
	fmt.Println("Football init")
}

func (f *Football) startPlay()  {
	fmt.Println("Football start Play ")
}

func (f *Football) endPlay()  {
	fmt.Println("Football end Play")
}

func NewFootball() *Football {
	f := new(Football)
	f.Game.init = f.init
	f.Game.startPlay = f.startPlay
	f.Game.endPlay = f.endPlay
	return f
}

func TestTemplate() {
	f := NewFootball()
	f.Play()
	cr := NewCricket()
	cr.Play()
}

func main() {
	TestTemplate()
}