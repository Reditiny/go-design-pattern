package visitor

// Visitor 访问者 定义了访问每个元素的方法 为每个具体元素提供一个访问操作的接口
// 不同的访问者实现该接口后即可以不同的方式访问元素
type Visitor interface {
	VisitMuseum(museum *Museum)
	VisitZoo(zoo *Zoo)
}
