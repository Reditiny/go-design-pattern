package interpreter

// AddExpression 非终结符表达式 - 加法
type AddExpression struct {
	Left  Expression
	Right Expression
}

func (ae *AddExpression) Interpret() int {
	return ae.Left.Interpret() + ae.Right.Interpret()
}
