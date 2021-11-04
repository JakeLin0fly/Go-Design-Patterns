package Factory_Pattern

import "fmt"

// 抽象产品
type IMouse interface {
	Use()
}
type IKeyboard interface {
	Use()
}

// 具体产品
type DellMouse struct {}
type DellKeyboard struct {}
type HPMouse struct {}
type HPKeyboard struct {}

func (d *DellMouse) Use()  {
	fmt.Println("user Dell mouse")
}
func (d *DellKeyboard) Use()  {
	fmt.Println("use Dell keyboard")
}

func (h *HPMouse) Use()  {
	fmt.Println("use HP mouse")
}
func (h *HPKeyboard) Use() {
	fmt.Println("use HP keyboard")
}

//// 抽象工厂
//type IFactory interface {
//	CreateMouse() IMouse
//	CreateKeyboard() IKeyboard
//}

// 具体工厂
// 戴尔工厂生产戴尔鼠标、戴尔键盘
type DellFactory struct {}
// 惠普工厂生产惠普鼠标、惠普键盘
type HPFactory struct {}

func (d *DellFactory) CreateMouse() IMouse {
	return &DellMouse{}
}
func (d *DellFactory) CreateKeyboard() IKeyboard {
	return &DellKeyboard{}
}

func (h *HPFactory) CreateMouse() IMouse {
	return &HPMouse{}
}
func (h *HPFactory) CreateKeyboard() IKeyboard {
	return &HPKeyboard{}
}