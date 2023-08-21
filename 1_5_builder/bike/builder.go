package bike

// Builder 自行车建造者接口 定义了产品的抽象接口，包含了构建产品的各个步骤
type Builder interface {
	BuildFrame()
	BuildSeat()
	BuildTyre()
	CreateBike() Bicycle
}
