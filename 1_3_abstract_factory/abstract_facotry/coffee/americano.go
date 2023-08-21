package coffee

// AmericanCoffee 具体咖啡 实现自己的具体功能
type AmericanCoffee struct {
}

func (a *AmericanCoffee) GetCoffeeName() string {
	return "美式咖啡"
}
