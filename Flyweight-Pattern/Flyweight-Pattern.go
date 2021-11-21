package Flyweight_Pattern

import "fmt"

/**
 * 俄罗斯方块中每次产生的方块只有几种形状（L、I、T、O）。如果每个不同的方块都是一个实例对象，
 * 这些对象就要占用很多的内存空间，此时适合利用享元模式进行实现。类图设计如下：
 * https://gitee.com/jakel-in/images/raw/master/2021-11/Flyweight-Pattern-Example.png
 */
// 抽象享元
type AbstractBox interface {
	GetShape() string
	Display(color string)
}

// 具体享元：I型
type IBox struct {}

func (i *IBox) GetShape() string {
	return "I"
}
func (i *IBox) Display(color string) {
	fmt.Println("Box Shape: ", i.GetShape(), ", Color: ", color)
}
// 具体享元：L型
type LBox struct {}

func (l *LBox) GetShape() string {
	return "L"
}
func (l *LBox) Display(color string) {
	fmt.Println("Box Shape: ", l.GetShape(), ", Color: ", color)
}
// 具体享元：L型
type TBox struct {}

func (t *TBox) GetShape() string {
	return "L"
}
func (t *TBox) Display(color string) {
	fmt.Println("Box Shape: ", t.GetShape(), ", Color: ", color)
}
// 具体享元：O型
type OBox struct {}

func (o *OBox) GetShape() string {
	return "I"
}
func (o *OBox) Display(color string) {
	fmt.Println("Box Shape: ", o.GetShape(), ", Color: ", color)
}

// 享元工厂：单例（饿汉）
func init() {
	factory = &boxFactory{hashMap: map[string]AbstractBox{
		"I": &IBox{},
		"L": &LBox{},
		"T": &TBox{},
		"O": &OBox{},
	}}
}
var factory *boxFactory
type boxFactory struct {
	hashMap map[string]AbstractBox
}

func (b *boxFactory) GetBox(shape string) AbstractBox {
	if val, ok:=  b.hashMap[shape]; ok {
		return val
	}
	return nil
}

func GetInstance() *boxFactory {
	return factory
}
