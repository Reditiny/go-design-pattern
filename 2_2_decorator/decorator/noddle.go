package decorator

// Noddle 现有对象 通过装饰器模式给其添加一些额外功能不需要改动已有代码
type Noddle struct {
}

func (e *Noddle) Price() int {
	return 10
}

func (e *Noddle) Content() string {
	return "一份炒面"
}
