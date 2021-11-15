package Proxy_Pattern

import "testing"

func TestProxy(t *testing.T) {
	// 测试：客户通过代理点购票
	proxy := ProxyPoint{}
	proxy.sell()
}
