package abstractFactory

import (
	"GoPractice/DesignPattern/CreationalPatterns/factory"
	"fmt"
)

const (
	SRed   = "Red"
	SBlue  = "Blue"
	SGreen = "Green"
)

type Color interface {
	fill(color string)
}

type Green struct {
}

func (g *Green) fill(color string) {
	fmt.Println("Fill color Blue")
}

type Blue struct {
}

func (b *Blue) fill(color string) {
	fmt.Println("Fill color Blue")
}

type Red struct {
}

func (r *Red) fill(color string) {
	fmt.Println("Fill color Red")
}

type FactoryAbstract interface {
	getShape(string) factory.Shape
	getColor(string) Color
}

type FactoryShape struct {
}

func (f *FactoryShape) getShape(strType string) factory.Shape {
	switch strType {
	case factory.SCircle:
		c := &factory.Circle{1}
		return c
	case factory.SSquare:
		return &factory.Square{
			1,
		}
	case factory.SRectangle:
		return &factory.Rectangle{
			3, 4,
		}
	default:
		return nil
	}
}

func (f *FactoryShape) getColor(strType string) Color {
	return nil
}

type FactoryColor struct {
}

func (f *FactoryColor) getShape(s string) factory.Shape {
	return nil
}

func (f *FactoryColor) getColor(s string) Color {
	switch s {
	case SRed:
		return &Red{}
	case SBlue:
		return &Blue{}
	case SGreen:
		return &Green{}
	default:
		return nil
	}
}
