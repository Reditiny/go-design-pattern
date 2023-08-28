package state

// Context 环境类
type Context struct {
	state   LiftState // 状态对象的引用，用于切换和使用不同的状态
	open    *Opening  // 其余为具体状态，封装在内部并不便于扩展新状态
	close   *Closing
	running *Running
	stop    *Stopping
}

func NewContext() *Context {
	c := new(Context)
	c.stop = new(Stopping)
	c.stop.context = c
	c.running = new(Running)
	c.running.context = c
	c.open = new(Opening)
	c.open.context = c
	c.close = new(Closing)
	c.close.context = c
	c.state = c.stop
	return c
}

func (c *Context) SetState(state LiftState) {
	c.state = state
}

func (c *Context) Open() {
	c.state.LiftOpen()
}

func (c *Context) Stop() {
	c.state.LiftStop()
}

func (c *Context) Run() {
	c.state.LiftRun()
}

func (c *Context) Close() {
	c.state.LiftClose()
}
