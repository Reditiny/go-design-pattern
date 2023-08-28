package state

import "fmt"

// Running 具体状态 封装了特定状态(运行)下的行为
type Running struct {
	context *Context
}

func (e *Running) LiftOpen() {
	fmt.Println("电梯运行时无法开启")
}

func (e *Running) LiftClose() {
	fmt.Println("电梯已经是关闭状态")
}

func (e *Running) LiftStop() {
	e.context.state = e.context.stop
	fmt.Println("电梯已停止")
}

func (e *Running) LiftRun() {
	fmt.Println("电梯已经是运行状态")
}
