package visitor

// Zoo 具体元素 实现了元素接口 提供接受访问者的具体实现
type Zoo struct {
	Name string
	Fee  int
}

func (z *Zoo) Accept(v Visitor) {
	v.VisitZoo(z)
}
