package Flyweight_Pattern

import (
	"fmt"
	"testing"
)

func TestFlyweightPattern(t *testing.T) {
	factory := GetInstance()

	iBox1 := factory.GetBox("I");
	iBox1.Display("Blue")

	lBox := factory.GetBox("L")
	lBox.Display("Red")

	iBox2 := factory.GetBox("I")
	iBox2.Display("Yellow")

	fmt.Println("IBox1 == lBox?   ", iBox1 == lBox)
	fmt.Println("IBox1 == iBox2?   ", iBox1 == iBox2)
}
