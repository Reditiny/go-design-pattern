package coffee

// LatteCoffee 具体咖啡 实现自己的具体功能
type LatteCoffee struct {
}

func (a *LatteCoffee) GetCoffeeName() string {
	return "拿铁咖啡"
}
