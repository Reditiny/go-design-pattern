package strategy

// B 具体策略 提供具体的算法实现或行为
type B struct{}

func (a *B) show() string {
	return "满减"
}
