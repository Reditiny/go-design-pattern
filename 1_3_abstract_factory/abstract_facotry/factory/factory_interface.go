package factory

import (
	"design_pattern/1_3_abstract_factory/abstract_facotry/coffee"
	"design_pattern/1_3_abstract_factory/abstract_facotry/dessert"
)

// IFactory 抽象工厂 生产咖啡与甜点 要增加产品时也要相应地增加工厂
type IFactory interface {
	CreateCoffee() coffee.ICoffee
	CreateDessert() dessert.Dessert
}
