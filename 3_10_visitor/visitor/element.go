package visitor

// Element  定义了一个接受访问者的接口
type Element interface {
	Accept(v Visitor)
}
