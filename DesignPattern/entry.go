package main

import "GoPractice/DesignPattern/CreationalPatterns/factory"

func main() {
	s := factory.ShapeFactory{}
	c := s.CreateShape("Circle")
	c.Draw()
	sq := s.CreateShape("Square")
	sq.Draw()
}
