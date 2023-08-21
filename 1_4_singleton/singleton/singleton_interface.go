package singleton

// Singleton 是单例模式接口
// 通过该接口可以避免 GetSomeInstance 返回一个包私有类型的指针
type Singleton interface {
	foo()
}
