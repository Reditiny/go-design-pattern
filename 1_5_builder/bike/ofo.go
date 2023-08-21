package bike

type OfoBuilder struct {
	Bicycle
}

func (b *OfoBuilder) BuildFrame() {
	b.Frame = "小黄车车架"
}

func (b *OfoBuilder) BuildSeat() {
	b.Seat = "小黄车车座"
}

func (b *OfoBuilder) BuildTyre() {
	b.Tyre = "小黄车车轮"
}

func (b *OfoBuilder) CreateBike() Bicycle {
	return b.Bicycle
}
