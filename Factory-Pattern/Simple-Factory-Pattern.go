package Factory_Pattern

import (
	"errors"
	"fmt"
)

// 抽象产品
type Shape interface {
	draw()
}

// 具体产品：圆形
type Circle struct { }

func (p *Circle) draw() {
	fmt.Println("draw a circle")
}

// 具体产品：三角形
type Rectangle struct { }

func (p *Rectangle) draw() {
	fmt.Println("draw a rectangle")
}

// 工厂
type Factory struct { }

func (p *Factory) create(name string) (Shape, error) {
	switch name {
	case "circle":
		return &Circle{}, nil
	case "rectangle":
		return &Rectangle{}, nil
	}
	return nil, errors.New("unsupported name")
}
