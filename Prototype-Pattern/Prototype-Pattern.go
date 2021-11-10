package Prototype_Pattern

import (
	"fmt"
	"strconv"
)

// 原型类，声明一个克隆自身的接口
type Prototype interface {
	Display()
	Clone() Prototype // 克隆接口
}

// 具体原型类，实现克隆自身的接口
type Resume struct {
	name string
	sex  string
	age  int64
}

func (r *Resume) Display() {
	fmt.Println("====== Resume ======")
	fmt.Println("\tname : " + r.name)
	fmt.Println("\tsex : " + r.sex)
	fmt.Println("\tage : " + strconv.FormatInt(r.age, 10))
}

func (r *Resume) Clone() Prototype {
	return &Resume{
		name: r.name,
		sex:  r.sex,
		age:  r.age,
	}
}
