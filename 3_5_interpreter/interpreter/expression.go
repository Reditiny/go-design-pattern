package interpreter

// Expression 抽象表达式接口
type Expression interface {
	Interpret() int
}
