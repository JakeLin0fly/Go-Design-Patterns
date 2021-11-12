package Adapter_Pattern

import "testing"

func TestAdapter(t *testing.T) {
	// 充电逻辑：插线板输出110V -> 接上适配器，110V转220V -> 充电输入电压220V
	phoneCharge := Adapter{adeptee: &Adaptee{}}
	phoneCharge.charge()
}
