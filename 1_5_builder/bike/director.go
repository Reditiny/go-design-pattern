package bike

// Director 指挥者类 负责使用 Builder 接口来构建产品，它定义了构建的顺序和流程。
type Director struct {
	Builder
}

func (d *Director) Construct() Bicycle {
	d.Builder.BuildFrame()
	d.Builder.BuildSeat()
	d.Builder.BuildTyre()
	return d.Builder.CreateBike()
}
