package Prototype_Pattern

import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	resume := Resume{
		name: "ZhangSan",
		sex:  "male",
		age:  18,
	}
	fmt.Println("print resume")
	resume.Display()
	fmt.Println()

	cp := resume.Clone()
	fmt.Println("print clone resume")
	cp.Display()
	fmt.Println()
}
