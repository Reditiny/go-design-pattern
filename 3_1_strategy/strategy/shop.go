package strategy

// Shop 环境类 持有一个策略类的引用 最终给客户端调用
type Shop struct {
	Strategy
}

func (s *Shop) SellStrategy() string {
	return s.Strategy.show()
}
