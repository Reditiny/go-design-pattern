package factory_method

type AmericanCoffeeFactory struct {
}

func (f *AmericanCoffeeFactory) CreateCoffee() ICoffee {
	return new(AmericanCoffee)
}
