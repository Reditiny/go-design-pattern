package flyweight

// Factory 享元工厂 提供复用的享元 可以改进为单例模式
type Factory struct {
	flyweights map[string]FlyWeight
}

func NewFactory() *Factory {
	factory := Factory{
		flyweights: make(map[string]FlyWeight, 4)}
	factory.flyweights["I"] = &IBox{}
	factory.flyweights["L"] = &LBox{}
	return &factory
}

func (f *Factory) GetFlyWeight(shape string) FlyWeight {
	// 省略对输入的判断
	return f.flyweights[shape]
}
