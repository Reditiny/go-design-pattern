package main

import (
	factory_method "design_pattern/1_2_facotry_method/facotry_method"
	"fmt"
)

func main() {
	coffeeFactory1 := new(factory_method.AmericanCoffeeFactory)
	coffee1 := coffeeFactory1.CreateCoffee()
	fmt.Println(coffee1.GetCoffeeName())
	coffeeFactory2 := new(factory_method.LatteCoffeeFactory)
	coffee2 := coffeeFactory2.CreateCoffee()
	fmt.Println(coffee2.GetCoffeeName())
}
