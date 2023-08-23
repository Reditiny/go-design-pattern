package composite

// MenuComponent 组件 抽像节点 定义了组合对象和叶子对象的共同接口使得客户端可以统一处理它们
type MenuComponent interface {
	Add(menuComponent *MenuComponent)
	Remove(menuComponent *MenuComponent)
	Print()
}
