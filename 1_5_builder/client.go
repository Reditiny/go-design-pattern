package main

import (
	"design_pattern/1_5_builder/bike"
	"design_pattern/1_5_builder/phone"
	"fmt"
)

func main() {
	// 自行车
	director1 := bike.Director{Builder: &bike.OfoBuilder{}}
	bike1 := director1.Construct()
	fmt.Println(bike1)
	director2 := bike.Director{Builder: &bike.MobileBuilder{}}
	bike2 := director2.Construct()
	fmt.Println(bike2)
	// 手机
	phoneBuilder := new(phone.Builder)
	phone1 := phoneBuilder.Cpu("骁龙").Memory("4GB").Screen("sum").Build()
	fmt.Println(phone1)
}
