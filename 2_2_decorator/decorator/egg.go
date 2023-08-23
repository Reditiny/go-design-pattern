package decorator

// Egg 具体装饰器
type Egg struct {
	FastFood
}

func (e *Egg) Price() int {
	return e.FastFood.Price() + 1
}

func (e *Egg) Content() string {
	return e.FastFood.Content() + "加个鸡蛋"
}
