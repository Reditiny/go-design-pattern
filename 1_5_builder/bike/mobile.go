package bike

// MobileBuilder 具体建造者 实现了 Builder 接口 负责实际构建产品的每个步骤并维护产品的内部状态
type MobileBuilder struct {
	Bicycle
}

func (b *MobileBuilder) BuildFrame() {
	b.Frame = "摩拜车架"
}

func (b *MobileBuilder) BuildSeat() {
	b.Seat = "摩拜车座"
}

func (b *MobileBuilder) BuildTyre() {
	b.Tyre = "摩拜车轮"
}

func (b *MobileBuilder) CreateBike() Bicycle {
	return b.Bicycle
}
