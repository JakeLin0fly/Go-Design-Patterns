package Bridge_Pattern

import "fmt"

/**
 * 现需设计实现一个视频播放器，有多个版本，分别支持不同的操作系统、解码不同的视频格式。
 * 如：Windows系统+AVI视频格式，Mac系统+RMVB视频格式等。
 * 类图：https://gitee.com/jakel-in/images/raw/master/2021-11/Bridge-Pattern-Example.png
 */

// 实现接口角色：解码器
type Decoder interface {
	decode(video string)
}

// 具体实现角色：AVI解码器
type AviDecoder struct{}

func (a *AviDecoder) decode(video string) {
	fmt.Print("decode AVI file : ", video)
}

// 具体实现角色：RMVB解码器
type RmvbDecoder struct{}

func (r *RmvbDecoder) decode(video string) {
	fmt.Print("decode RMVB file : ", video)
}

// 抽象角色：操作系统
type OperatingSystem struct {
	decoder Decoder
}

func (a *OperatingSystem) play(video string) {
	fmt.Println("Play video : ", video)
}

// 扩充抽象角色：具体操作系统 Windows
type Windows struct {
	OperatingSystem
}

func (w *Windows) play(video string) {
	fmt.Print("Play on Windows, ")
	w.decoder.decode(video)
	fmt.Println()
}

// 扩充抽象角色：具体操作系统 Mac
type Mac struct {
	OperatingSystem
}

func (m *Mac) play(video string) {
	fmt.Print("Play on Mac, ")
	m.decoder.decode(video)
	fmt.Println()
}
