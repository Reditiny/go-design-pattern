package factory_method

// CoffeeFactory 抽象工厂 生产咖啡 要增加产品时也要相应地增加工厂
type CoffeeFactory interface {
	CreateCoffee() ICoffee
}
