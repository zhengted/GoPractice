### 概念
- 命令模式是一种数据驱动的设计模式。请求以命令的形式包裹在对象中，并传给调用对象。调用对象寻找可以处理该命令的合适的对象，并把该命令传给相应的对象
- **命令模式就是面向对象化的回调**
- 角色说明
    - receiver 接收者，最终使用命令的对象
    - Command 命令基类，所有的命令派生类都用一个execute，确保封装性
    - invoker 发起命令的角色
- **优点**：降低耦合度，命令的发起者只需要知道发起的是什么命令将对应的接收者作为参数传入。增加命令比较简单
- **好像没什么缺点**

- 代码
    - 以下代码以电视机为例，实现三种命令，开机关机，切换频道。该实例没有发起者 只有接收者和命令相关定义

```Golang
package main

import "fmt"

type TV struct {}

func (t TV) Open()  {
	fmt.Println("Open TV\n")
}

func (t TV) Close() {
	fmt.Println("Close TV\n")
}

func (t TV) ChangeChanel()  {
	fmt.Println("ChangeChannel\n")
}

// 命令基类 及 对应派生类
type Command interface {
	execute()
}

type OpenCommand struct {
	receiver TV
}


func (r *OpenCommand) execute()  {
	fmt.Println("Command Open TV\n")
	r.receiver.Open()
}

type CloseCommand struct {
	receiver TV
}

func (r *CloseCommand) execute()  {
	fmt.Println("Command Close TV\n")
	r.receiver.Close()
}

type ChangeChannelCommand struct {
	receiver TV
}

func (r *ChangeChannelCommand) execute()  {
	fmt.Println("Command Change channel\n")
	r.receiver.ChangeChanel()
}

func TestCommandPattern() {
	t := TV{}
	c1 := OpenCommand{receiver: t}
	c2 := CloseCommand{receiver: t}
	c3 := ChangeChannelCommand{receiver: t}
	c1.execute()
	c3.execute()
	c2.execute()
}
```