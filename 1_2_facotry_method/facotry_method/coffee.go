package factory_method

import "fmt"

// Coffee 抽象咖啡 实现咖啡接口的公共部分 也可以没有这个结构并将下面两个函数在具体咖啡中实现
type Coffee struct {
}

func (c *Coffee) AddSugar() {
	fmt.Println("加糖")
}

func (c *Coffee) AddMilk() {
	fmt.Println("加奶")
}
