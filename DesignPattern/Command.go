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