package strategy

// A 具体策略 提供具体的算法实现或行为
type A struct{}

func (a *A) show() string {
	return "折扣"
}
