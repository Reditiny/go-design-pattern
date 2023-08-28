package state

// LiftState 状态接口 其中包含了具体状态类需要实现的方法
type LiftState interface {
	LiftOpen()
	LiftClose()
	LiftStop()
	LiftRun()
}
