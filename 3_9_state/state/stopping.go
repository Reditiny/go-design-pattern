package state

import "fmt"

// Stopping 具体状态 封装了特定状态(停止)下的行为
type Stopping struct {
	context *Context
}

func (e *Stopping) LiftOpen() {
	e.context.state = e.context.open
	fmt.Println("电梯已开启")
}

func (e *Stopping) LiftClose() {
	e.context.state = e.context.close
	fmt.Println("电梯已关闭")
}

func (e *Stopping) LiftStop() {
	fmt.Println("电梯已经是停止状态")
}

func (e *Stopping) LiftRun() {
	fmt.Println("电梯关闭时才可运行")
}
