package command

import "fmt"

// OrderCommand 具体命令 实现了命令接口，将请求（DinnerOrder）绑定到接收者（Receiver）并执行相应的操作。
type OrderCommand struct {
	Receiver    Chief
	DinnerOrder Order
}

func (c *OrderCommand) execute() {
	fmt.Printf("给%d号桌备餐\n", c.DinnerOrder.TableId)
	for key, value := range c.DinnerOrder.FoodDir {
		c.Receiver.MakeFood(key, value)
	}
}
