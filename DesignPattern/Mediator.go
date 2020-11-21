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