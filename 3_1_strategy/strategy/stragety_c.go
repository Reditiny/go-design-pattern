package strategy

// C 具体策略 提供具体的算法实现或行为
type C struct{}

func (a *C) show() string {
	return "赠送"
}
