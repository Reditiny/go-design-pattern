package factory

import (
	"design_pattern/1_3_abstract_factory/abstract_facotry/coffee"
	"design_pattern/1_3_abstract_factory/abstract_facotry/dessert"
)

type AmericanFactory struct {
}

func (f *AmericanFactory) CreateCoffee() coffee.ICoffee {
	return new(coffee.AmericanCoffee)
}

func (f *AmericanFactory) CreateDessert() dessert.Dessert {
	return new(dessert.Mousse)
}
