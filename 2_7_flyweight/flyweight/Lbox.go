package flyweight

type LBox struct {
}

func (l *LBox) GetShape() string {
	return "L"
}
