package simple_facotry

// AmericanCoffee 具体咖啡 实现自己的具体功能
type AmericanCoffee struct {
	Coffee
}

func (a *AmericanCoffee) GetCoffeeName() string {
	return "美式咖啡"
}
