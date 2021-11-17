package Decorator_Pattern

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	originCoffee := ShortBlack{}
	fmt.Println(originCoffee.description(), " : ", originCoffee.cost())

	addSugar := CoffeeWithSugar{
		CoffeeDecorator{
			coffee: &originCoffee,
		},
	}
	fmt.Println(addSugar.description(), " : ", addSugar.cost())

	addSugarMilk := CoffeeWithMilk{
		CoffeeDecorator{
			coffee: &addSugar,
		},
	}
	fmt.Println(addSugarMilk.description(), " : ", addSugarMilk.cost())
}
