package main

import (
	"design_pattern/1_3_abstract_factory/abstract_facotry/factory"
	"fmt"
)

func main() {
	americanFactory := new(factory.AmericanFactory)
	coffee1 := americanFactory.CreateCoffee()
	fmt.Println(coffee1.GetCoffeeName())
	dessert1 := americanFactory.CreateDessert()
	fmt.Println(dessert1.GetDessertName())
	italyFactory := new(factory.ItalyFactory)
	coffee2 := italyFactory.CreateCoffee()
	fmt.Println(coffee2.GetCoffeeName())
	dessert2 := italyFactory.CreateDessert()
	fmt.Println(dessert2.GetDessertName())
}
