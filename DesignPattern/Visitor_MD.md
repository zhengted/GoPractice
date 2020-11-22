### 访问者模式
- 概念：我们使用了一个访问者类，它改变了元素类的执行算法。通过这种方式，元素的执行算法可以随着访问者改变而改变。这种类型的设计模式属于行为型模式。根据模式，元素对象已接受访问者对象，这样访问者对象就可以处理元素对象上的操作。
- 缺点极其明显的一种设计模式
- 违反了迪米特法则 暴露了对象的内部细节
- 违反了依赖倒转原则 依赖了具体类而不是抽象类
- 优点
    -  1、符合单一职责原则。 2、优秀的扩展性。 3、灵活性。

```golang
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
```