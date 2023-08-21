package simple_facotry

// SimpleCoffeeFactory 简单咖啡工厂 生产咖啡 调用者通过方法来获取产品
type SimpleCoffeeFactory struct {
}

func (f *SimpleCoffeeFactory) CreateCoffee(coffeeName string) ICoffee {
	var coffee ICoffee
	if coffeeName == "美式咖啡" {
		coffee = new(AmericanCoffee)
	} else if coffeeName == "拿铁咖啡" {
		coffee = new(LatteCoffee)
	} else {
		panic("没有此类咖啡")
	}
	return coffee
}
