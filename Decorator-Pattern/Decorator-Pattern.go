package Decorator_Pattern

/**
 * 现在有一个咖啡店售卖多种单品咖啡（Espresso、ShortBlack等）销售火爆，
 * 店家推出组合咖啡（单品咖啡+配料），组合很多，如：单品咖啡+牛奶，单品咖啡+糖 等等。
 * 现需要更新系统支持组合咖啡价格计算。
 * 类图：https://gitee.com/jakel-in/images/raw/master/2021-11/Decorator-Pattern-example.jpg
 */

// 【抽象构件】 咖啡抽象接口
type Coffee interface {
	description() string
	cost() float64
}

// 【具体构件】 咖啡单品类：Espresso
type Espresso struct{}

func (e *Espresso) description() string {
	return "Espresso(12)"
}
func (e *Espresso) cost() float64 {
	return 12
}

// 【具体构件】 咖啡单品类：ShortBlack
type ShortBlack struct{}

func (s *ShortBlack) description() string {
	return "ShortBlack(13)"
}
func (s *ShortBlack) cost() float64 {
	return 13
}

// 【抽象装饰】 组合咖啡基类
type CoffeeDecorator struct {
	coffee Coffee
}

func (c *CoffeeDecorator) description() string {
	return c.coffee.description()
}
func (c *CoffeeDecorator) cost() float64 {
	return c.coffee.cost()
}

// 【具体装饰】 加糖
type CoffeeWithSugar struct {
	CoffeeDecorator // 继承
}

func (c *CoffeeWithSugar) description() string {
	return c.CoffeeDecorator.description() + " + Sugar(2)"
}
func (c *CoffeeWithSugar) cost() float64 {
	return c.CoffeeDecorator.cost() + 2
}

// 【具体装饰】 加牛奶
type CoffeeWithMilk struct {
	CoffeeDecorator // 继承
}

func (c *CoffeeWithMilk) description() string {
	return c.CoffeeDecorator.description() + " + Sugar(3)"
}
func (c *CoffeeWithMilk) cost() float64 {
	return c.CoffeeDecorator.cost() + 3
}
