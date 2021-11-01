package Singleton_Pattern

func init() {
	// 恶汉模式
	hInstance = &hungerSingleton{}
}
type hungerSingleton struct { }

var hInstance *hungerSingleton

func GetHungerInstance() *hungerSingleton {
	return hInstance
}