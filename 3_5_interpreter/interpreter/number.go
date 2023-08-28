package interpreter

// NumberExpression 终结符表达式
type NumberExpression struct {
	Number int
}

func (ne *NumberExpression) Interpret() int {
	return ne.Number
}
