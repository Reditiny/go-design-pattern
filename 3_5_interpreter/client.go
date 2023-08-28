package main

import (
	"design_pattern/3_5_interpreter/interpreter"
	"fmt"
)

func main() {
	// 构建表达式
	expression1 := &interpreter.NumberExpression{Number: 10}
	expression2 := &interpreter.NumberExpression{Number: 3}
	expression3 := &interpreter.AddExpression{Left: expression1, Right: expression2}
	expression4 := &interpreter.NumberExpression{Number: 9}
	expression5 := interpreter.SubtractExpression{Left: expression3, Right: expression4}
	// 计算结果
	fmt.Println(expression5.Interpret())

}
