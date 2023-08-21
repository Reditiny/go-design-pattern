package factory

import (
	"design_pattern/1_3_abstract_factory/abstract_facotry/coffee"
	"design_pattern/1_3_abstract_factory/abstract_facotry/dessert"
)

type ItalyFactory struct {
}

func (f *ItalyFactory) CreateCoffee() coffee.ICoffee {
	return new(coffee.LatteCoffee)
}

func (f *ItalyFactory) CreateDessert() dessert.Dessert {
	return new(dessert.Tiramisu)
}
