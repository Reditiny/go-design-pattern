package interpreter

// SubtractExpression 非终结符表达式 - 减法
type SubtractExpression struct {
	Left  Expression
	Right Expression
}

func (se *SubtractExpression) Interpret() int {
	return se.Left.Interpret() - se.Right.Interpret()
}
