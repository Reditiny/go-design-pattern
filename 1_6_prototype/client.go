package main

import (
	"design_pattern/1_6_prototype/prototype"
	"fmt"
	"time"
)

func main() {
	human1 := prototype.Human{Flesh: "血肉之躯"}
	soul := make([]time.Time, 2, 2)
	soul[0] = time.Now()
	soul[1] = time.Unix(123, 456)
	human1.Soul = soul
	// 原型复制 由于返回值均为接口类型 因此要做类型断言
	human2, ok := human1.Clone().(*prototype.Human)
	if ok { // 验证结构内的切片是否为深拷贝
		fmt.Println(human1.Soul)
		fmt.Println(human2.Soul)
		human1.Soul[1] = time.Now()
		fmt.Println(human1.Soul)
		fmt.Println(human2.Soul)
	}

}
