package Factory_Pattern

import (
	"errors"
	"fmt"
)

// 抽象产品
type IShape interface {
	Draw()
}

// 具体产品：圆形
type Circle struct { }

func (p *Circle) Draw() {
	fmt.Println("draw a circle")
}

// 具体产品：三角形
type Rectangle struct { }

func (p *Rectangle) Draw() {
	fmt.Println("draw a rectangle")
}

// 工厂
type Factory struct { }

func (p *Factory) create(name string) (IShape, error) {
	switch name {
	case "circle":
		return &Circle{}, nil
	case "rectangle":
		return &Rectangle{}, nil
	}
	return nil, errors.New("unsupported shape")
}
