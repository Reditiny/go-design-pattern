package main

import (
	"design_pattern/3_1_strategy/strategy"
	"fmt"
)

func main() {
	shop := strategy.Shop{}
	s1 := &strategy.A{}
	s2 := &strategy.B{}
	s3 := &strategy.C{}
	// 活动1
	shop.Strategy = s1
	fmt.Println(shop.SellStrategy())
	// 活动2
	shop.Strategy = s2
	fmt.Println(shop.SellStrategy())
	// 活动3
	shop.Strategy = s3
	fmt.Println(shop.SellStrategy())
}
