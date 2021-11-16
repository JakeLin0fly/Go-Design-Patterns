package Adapter_Pattern

import (
	"fmt"
	"testing"
)

func TestObjectAdapter(t *testing.T) {
	// 电脑A：未接适配器
	computer_A := Computer{
		sdCard: &SDCard{},
	}
	// 电脑B：接上适配器
	computer_B := Computer{
		sdCard: &AdapterSD_TF{
			tfCard: &TFCard{},
		},
	}

	fmt.Println("============= test reading =============")
	fmt.Println("Computer A(no Adapter): \n\t", computer_A.readSD())
	fmt.Println("Computer B(with Adapter): \n\t", computer_B.readSD())

	fmt.Println("\n============= test writing =============")
	fmt.Print("Computer A(no Adapter): \n\t")
	computer_A.writeSD("test writing")

	fmt.Print("Computer B(with Adapter): \n\t")
	computer_B.writeSD("test writing")
}
