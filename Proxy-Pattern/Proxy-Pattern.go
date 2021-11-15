package Proxy_Pattern

import "fmt"
/**
 * 以前交通、互联网没那么发达，购买火车票可以在火车站直接购买，但通常人们是在车票代售点进行购买。
 * 这个例子中火车站是目标对象，代售点是代理对象。
 */

// 抽象主题（买票接口）
type SellTickets interface {
	sell()
}

// 具体主题（代理目标）：火车站
type TrainStation struct {}

func (t *TrainStation) sell() {
	fmt.Println("sell ticket from train station")
}

// 代理：代售点
type ProxyPoint struct {
	station TrainStation
}

func (p *ProxyPoint) sell()  {
	fmt.Print("proxy point charges handling fees, ")
	p.station.sell()
}