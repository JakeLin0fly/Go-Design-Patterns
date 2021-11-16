package Adapter_Pattern

import "fmt"

/**
 * ## 对象适配器示例 ##
 * 现有电脑仅有SD卡读写接口，而数据存储使用TF卡，
 * 此时就需要一个适配器（转接器），使得能通过电脑上仅有的SD卡接口实现对TF卡的读写。
 * 类图：https://gitee.com/jakel-in/images/raw/master/2021-11/Object-Adapter-Pattern.jpg
 */

// 卡读写接口
type Card interface {
	read() string
	write(msg string)
}

//SD卡实现类
type SDCard struct{}

func (s *SDCard) read() string {
	return "SD card read a msg: hello word SD"
}
func (s *SDCard) write(msg string) {
	fmt.Println("SD card write msg : ", msg)
}

// TF卡实现类
type TFCard struct{}

func (s *TFCard) read() string {
	return "TF card read a msg: hello word SD"
}
func (s *TFCard) write(msg string) {
	fmt.Println("TF card write msg : ", msg)
}

// 电脑类，仅支持SD卡
type Computer struct {
	sdCard Card
}

func (c *Computer) readSD() string {
	return c.sdCard.read()
}

func (c *Computer) writeSD(msg string) {
	c.sdCard.write(msg)
}

// 定义适配器类（SD兼容TF）
type AdapterSD_TF struct {
	SDCard
	tfCard *TFCard
}

func (a *AdapterSD_TF) read() string {
	return "adapter read TF card, " + a.tfCard.read()
}

func (a *AdapterSD_TF) write(msg string) {
	fmt.Print("adapter write TF card, ")
	a.tfCard.write(msg)
}
