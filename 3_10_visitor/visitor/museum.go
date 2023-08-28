package visitor

// Museum 具体元素 实现了元素接口 提供接受访问者的具体实现
type Museum struct {
	Name string
	Fee  int
}

func (m *Museum) Accept(v Visitor) {
	v.VisitMuseum(m)
}
