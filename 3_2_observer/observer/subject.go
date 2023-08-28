package observer

// Subject 抽象主题 提供增删订阅者和发布接口
type Subject interface {
	Attach(observer *Observer)
	Detach(observer *Observer)
	Notify(message string)
}
