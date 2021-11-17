package Facade_Pattern

import "fmt"

/**
 * 小明每天作息很规律，其中包括以下场景操作：
 *  下班回家：开灯、打开空调、打开电视；
 *  睡觉：关电视、关灯；
 * 小明购买了一个天猫精灵，通过语言控制这些设备开启关闭。设计类图如下：
 * https://gitee.com/jakel-in/images/raw/master/2021-11/Facade-Pattern-Example.png
 */

// 子系统角色：灯
type Light struct{}

func (l *Light) on() {
	fmt.Print("Light(on)")
}
func (l *Light) off() {
	fmt.Print("Light(off)")
}

// 子系统角色：电视
type TV struct{}

func (t *TV) on() {
	fmt.Print("TV(on)")
}
func (t *TV) off() {
	fmt.Print("TV(off)")
}

// 子系统角色：空调
type AirCondition struct{}

func (a *AirCondition) on() {
	fmt.Print("AirCondition(on)")
}
func (a *AirCondition) off() {
	fmt.Print("AirCondition(off)")
}

// 外观角色：天猫精灵
type TmallGenie struct {
	light        *Light
	tv           *TV
	airCondition *AirCondition
}

func (t *TmallGenie) Say(scene string) {
	if "go home" == scene {
		t.goHome()
		return
	}
	if "sleep" == scene {
		t.sleep()
	}
}

func (t *TmallGenie) goHome() {
	fmt.Print("scene: go home. operations: ")
	t.light.on()
	fmt.Print(" + ")
	t.airCondition.on()
	fmt.Print(" + ")
	t.tv.on()
	fmt.Println()
}

func (t *TmallGenie) sleep() {
	fmt.Print("scene: sleep. operations: ")
	t.tv.off()
	fmt.Print(" + ")
	t.light.off()
	fmt.Println()
}
