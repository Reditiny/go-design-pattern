package main

import (
	"design_pattern/1_1_simple_facoty/simple_facotry"
	"fmt"
)

func main() {
	coffeeFactory := new(simple_facotry.SimpleCoffeeFactory)
	coffee1 := coffeeFactory.CreateCoffee("美式咖啡")
	fmt.Println(coffee1.GetCoffeeName())
	coffee2 := coffeeFactory.CreateCoffee("拿铁咖啡")
	fmt.Println(coffee2.GetCoffeeName())
	coffee3 := coffeeFactory.CreateCoffee("中国咖啡")
	fmt.Println(coffee3.GetCoffeeName())
}
