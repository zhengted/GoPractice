package factory

import "fmt"

const (
	SCircle    = "Circle"
	SRectangle = "Rectangle"
	SSquare    = "Square"
)

type Shape interface {
	Draw()
}

type Circle struct {
	Radius int
}

func (c *Circle) Draw() {
	fmt.Println("draw a circle")
}

type Square struct {
	Edge int
}

func (s Square) Draw() {
	fmt.Println("draw a square")
}

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) Draw() {
	fmt.Println("draw a Rectangle")
}

type ShapeFactory struct {
}

func (s ShapeFactory) CreateShape(shapeType string) Shape {
	switch shapeType {
	case SCircle:
		c := &Circle{1}
		return c
	case SSquare:
		return &Square{
			1,
		}
	case SRectangle:
		return &Rectangle{
			3, 4,
		}
	default:
		return nil
	}
}
