package Factory_Pattern

//// 抽象产品
//type IShape interface {
//	Draw()
//}
//
//// 具体产品：圆形
//type Circle struct { }
//
//func (p *Circle) Draw() {
//	fmt.Println("draw a circle")
//}
//
//// 具体产品：长方形
//type Rectangle struct { }
//
//func (p *Rectangle) Draw() {
//	fmt.Println("draw a rectangle")
//}

// 抽象工厂
type IFactory interface {
	CreateShape()
}

// 具体工厂：圆形工厂
type CircleFactory struct {}

func (p *CircleFactory) CreateShape() IShape {
	return &Circle{}
}

// 具体工厂：长方形工厂
type RectangleFactory struct {}

func (p *RectangleFactory) CreateShape() IShape {
	return &Rectangle{}
}