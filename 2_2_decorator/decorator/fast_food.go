package decorator

// FastFood 抽象装饰器
type FastFood interface {
	Price() int
	Content() string
}
