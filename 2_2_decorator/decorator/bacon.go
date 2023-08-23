package decorator

// Bacon 具体装饰器
type Bacon struct {
	FastFood
}

func (e *Bacon) Price() int {
	return e.FastFood.Price() + 3
}

func (e *Bacon) Content() string {
	return e.FastFood.Content() + "加个培根"
}
