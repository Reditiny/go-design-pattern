package flyweight

type IBox struct {
}

func (i *IBox) GetShape() string {
	return "L"
}
