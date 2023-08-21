package main

import (
	"design_pattern/1_4_singleton/singleton"
	"fmt"
)

func main() {
	eagerSingleton1 := singleton.GetEagerSingleton()
	eagerSingleton2 := singleton.GetEagerSingleton()
	fmt.Printf("是否为同一个实例 ? %v\n", eagerSingleton1 == eagerSingleton2)
	lazySingleton1 := singleton.GetLazySingleton()
	lazySingleton2 := singleton.GetLazySingleton()
	fmt.Printf("是否为同一个实例 ? %v\n", lazySingleton1 == lazySingleton2)
}
