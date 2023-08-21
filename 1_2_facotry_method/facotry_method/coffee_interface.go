package factory_method

// ICoffee 咖啡接口 定义规范 描述主要特性和功能
type ICoffee interface {
	GetCoffeeName() string
	AddSugar()
	AddMilk()
}
