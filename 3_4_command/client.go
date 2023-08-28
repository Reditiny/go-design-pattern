package main

import (
	"design_pattern/3_4_command/command"
)

func main() {
	// 餐桌下单
	mealMap := make(map[string]int, 2)
	mealMap["鱼香肉丝"] = 1
	mealMap["青椒炒蛋"] = 1
	mealMap["米饭"] = 2
	order := command.Order{TableId: 1, FoodDir: mealMap}
	// 服务员发出命令
	var com command.Command
	com = &command.OrderCommand{Receiver: command.Chief{}, DinnerOrder: order}
	waiter := command.Waiter{Command: &com}
	waiter.MakeCommand()

}
