package state

import "fmt"

// Opening 具体状态 封装了特定状态(开启)下的行为
type Opening struct {
	context *Context
}

func (e *Opening) LiftOpen() {
	fmt.Println("电梯已经是开启状态")
}

func (e *Opening) LiftClose() {
	e.context.state = e.context.close
	fmt.Println("电梯已关闭")
}

func (e *Opening) LiftStop() {
	fmt.Println("电梯已经是停止状态")
}

func (e *Opening) LiftRun() {
	fmt.Println("电梯开启状态无法运行，请先关闭")
}
