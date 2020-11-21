### 状态模式
- 先贴个菜鸟教程的[链接](https://www.runoob.com/design-pattern/state-pattern.html)
- 状态模式：
	- 一个容易和命令模式混淆的模式，两者的意图都是为了消除冗余的If语句
	- 当代码中含有大量和对象状态相关的条件语句，且对象的行为依赖于它的状态
- 区别：
	- 命令模式的接口中只有一个方法，（命令封装成类 类内部的执行方法）
	- 状态模式的接口中有一个或者多个方法
	- 个人理解：这两个模式的区别是封装了什么内容
	    - 命令模式：封装了请求
	    - 状态模式：封装了状态转换的过程（一定程度上有点像观察者）
- 游戏开发过程中能看到的
    - 以Ai状态机为例，怪物的归属状态有很多种：出生归属（谁给了第一刀归谁），伤害归属（所在团或队伍伤害最高，归属者为团长或队长）
    - 归属的模式通过配置进行修改
    - 代码中也可以修改其对应的状态，但是无需关心如何转换的，以及什么状态下对应的Action
- 代码
```golang
package main

import "fmt"

/*
	状态模式：
		一个容易和命令模式混淆的模式，两者的意图都是为了消除冗余的If语句
		当代码中含有大量和对象状态相关的条件语句，且对象的行为依赖于它的状态
	区别：
		命令模式的接口中只有一个方法，（命令封装成类 类内部的执行方法）
		状态模式的接口中有一个或者多个方法

*/
type Context struct {
	state IState
}

func (c *Context) setState(state IState)  {
	c.state = state
}

func (c *Context) getState() IState {
	return c.state
}

type IState interface {
	doAction(context *Context)
}

type StartState struct {}

func (ss *StartState) doAction(ctx *Context)  {
	fmt.Println("Player is in start state")
	ctx.setState(ss)
}

type StopState struct {}

func (s *StopState) doAction(ctx *Context)  {
	fmt.Println("Player is in stop state")
	ctx.setState(s)
}

// 实际开发中 doAction这种操作是单独一个线程执行的
// 以游戏AI举例  怪物的归属模式有：出生归属模式，伤害归属模式 这种模式其实都记录在怪物的状态中
// 如果因为特殊需求修改了配置表中的归属模式会导致怪物的AI寻敌方式发生改变
func TestState()  {
	ctx := new(Context)
	start := new(StartState)
	start.doAction(ctx)
	stop := new(StopState)
	stop.doAction(ctx)
}


```