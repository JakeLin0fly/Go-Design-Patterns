package Adapter_Pattern

import "fmt"

// 目标
type Target interface {
	charge()
}

// 需适配类/接口
type Adaptee struct{}

func (a *Adaptee) SpecialCharge() {
	fmt.Println("Patch board output 110V.")
}

// 对象适配器
type Adapter struct {
	adeptee *Adaptee
}

func (a Adapter) charge() {
	// 转换
	fmt.Println("Connect the adapter, 110V to 220V.")
	a.adeptee.SpecialCharge()
}
