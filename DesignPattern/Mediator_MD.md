### 中介者模式
- 本人才疏学浅 看着[菜鸟教程的中介者模式](https://www.runoob.com/design-pattern/mediator-pattern.html)敲的
- 最常见的应用就是MVC
- 没有Web开发经验 不太了解MVC具体的写法
- 只能仿照着页面上的聊天室的写法写了个聊天

- 代码
```Golang
package main

import "fmt"

/*
	中介者模式
	中介者模式（Mediator Pattern）是用来降低多个对象和类之间的通信复杂性。
	这种模式提供了一个中介类，该类通常处理不同类之间的通信，并支持松耦合，使代码易于维护。
	中介者模式属于行为型模式。
	举例：MVC框架
*/

type ChatRoom struct {

}

func (c ChatRoom) showMessage(user *User,message string)  {
	fmt.Printf("%s:%s\n",user.name,message)
}

type User struct {
	name string
	chat ChatRoom
}

func (u *User) User(name string,chat ChatRoom) {
	u.name = name
	u.chat = chat
}

func (u *User) sendMessage(message string)  {
	u.chat.showMessage(u,message)
}

func TestMediator() {
	c := ChatRoom{}
	john := User{
		chat:c,
		name: "john",
	}
	robert := User{
		chat: c,
		name: "robert",
	}
	john.sendMessage("Hello Robert")
	robert.sendMessage("Hi,John")
}

func main() {
	TestMediator()
}
```