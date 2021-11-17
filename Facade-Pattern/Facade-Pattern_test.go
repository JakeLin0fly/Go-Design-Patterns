package Facade_Pattern

import "testing"

// 测试：模拟客户场景操作
func TestFacadePattern(t *testing.T) {
	tmallGenie := TmallGenie{
		light:        &Light{},
		tv:           &TV{},
		airCondition: &AirCondition{},
	}
	tmallGenie.Say("sleep")
	tmallGenie.Say("go home")
}
