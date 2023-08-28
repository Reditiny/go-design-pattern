package state

import "fmt"

// Closing 具体状态 封装了特定状态(关闭)下的行为
type Closing struct {
	context *Context
}

func (e *Closing) LiftOpen() {
	fmt.Println("停止状态下才能打开电梯")
}

func (e *Closing) LiftClose() {
	fmt.Println("电梯已经是关闭状态")
}

func (e *Closing) LiftStop() {
	e.context.state = e.context.stop
	fmt.Println("电梯已停止")
}

func (e *Closing) LiftRun() {
	e.context.state = e.context.running
	fmt.Println("电梯已运行")
}
