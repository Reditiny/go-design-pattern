package strategy

// Strategy 策略接口 给出具体策略类所需的接口
type Strategy interface {
	show() string // 可以将具体算法封装为私有 外界不可见
}
