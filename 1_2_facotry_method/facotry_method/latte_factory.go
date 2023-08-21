package factory_method

type LatteCoffeeFactory struct {
}

func (f *LatteCoffeeFactory) CreateCoffee() ICoffee {
	return new(LatteCoffee)
}
