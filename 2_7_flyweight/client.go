package main

import (
	"design_pattern/2_7_flyweight/flyweight"
	"fmt"
)

func main() {
	factory := flyweight.NewFactory()
	flyWeight1 := factory.GetFlyWeight("I")
	flyWeight2 := factory.GetFlyWeight("I")
	// 两次得到的是同一个对象
	fmt.Println(flyWeight1 == flyWeight2)
	flyWeight3 := factory.GetFlyWeight("L")
	fmt.Println(flyWeight3.GetShape())
}
