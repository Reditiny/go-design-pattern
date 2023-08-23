package main

import (
	"design_pattern/2_2_decorator/decorator"
	"fmt"
)

func main() {
	var food decorator.FastFood
	// 点个炒饭
	food = &decorator.Noddle{}
	// 加个蛋
	food = &decorator.Egg{FastFood: food}
	// 加个培根
	food = &decorator.Bacon{FastFood: food}
	fmt.Printf("%s 一共 %d 元\n", food.Content(), food.Price())

	// 点个炒饭
	food = &decorator.Noddle{}
	// 加个蛋
	food = &decorator.Egg{FastFood: food}
	// 再加个蛋
	food = &decorator.Egg{FastFood: food}
	fmt.Printf("%s 一共 %d 元\n", food.Content(), food.Price())
}
