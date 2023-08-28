package responsibility_chan

// Handler 抽象处理者
type Handler interface {
	Handle(days int)
}
