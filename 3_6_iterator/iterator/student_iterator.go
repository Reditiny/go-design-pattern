package iterator

// StudentIterator 抽象迭代器 定义访问和遍历聚合元素的接口，通常包含 HasNext Next 等方法。
type StudentIterator interface {
	HasNext() bool
	Next() *Student
}
