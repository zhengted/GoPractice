package main

import "fmt"

/*
	访问者模式是一个缺点非常明显的模式
		1.违反了迪米特法则 向访问者透露了类的实现细节
		2.违反了依赖倒置原则，依赖了具体类而不是抽象类
*/

type ComputerPart interface {
	accept(cpv ComputerPartVisitor)
}

type Keyboard struct {}

func (k *Keyboard) accept(cpv ComputerPartVisitor)  {
	cpv.visitKb(k)
}

type Monitor struct {}

func (m *Monitor) accept(cpv ComputerPartVisitor)  {
	cpv.visitMon(m)
}

type Mouse struct {}

func (m *Mouse) accept(cpv ComputerPartVisitor)  {
	cpv.visitMou(m)
}

type Computer struct {
	parts []ComputerPart
}

func (c *Computer) Computer() *Computer {
	c.parts = append(c.parts,new(Mouse))
	c.parts = append(c.parts,new(Keyboard))
	c.parts = append(c.parts,new(Monitor))
	return c
}

func (c *Computer) accept(cpv ComputerPartVisitor)  {
	for _,v := range c.parts {
		v.accept(cpv)
	}
	cpv.visitCp(c)
}

type ComputerPartVisitor interface {
	visitCp(computer *Computer)
	visitKb(k *Keyboard)
	visitMon(m *Monitor)
	visitMou(m *Mouse)
}

type ConcreteCPV struct {

}

func (c *ConcreteCPV) visitCp(computer *Computer)  {
	fmt.Println("Visit Computer")
}

func (c *ConcreteCPV) visitKb(k *Keyboard)  {
	fmt.Println("Visit KeyBoard")
}

func (c *ConcreteCPV) visitMon(m *Monitor)  {
	fmt.Println("Visit Monitor")
}

func (c *ConcreteCPV) visitMou(m *Mouse)  {
	fmt.Println("Visit Mouse")
}

func TestVisitor()  {
	cp := new(Computer).Computer()
	cp.accept(new(ConcreteCPV))
}

func main()  {
	TestVisitor()
}