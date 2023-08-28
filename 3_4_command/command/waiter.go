package command

// Waiter 调用者 负责触发命令并确保命令被执行，不需要关心命令的具体细节
type Waiter struct {
	Command *Command
}

func (w *Waiter) MakeCommand() {
	(*(w.Command)).execute()
}
