package singleton

// 饿汉式单例 包私有
type eager struct {
	data string
}

func (e *eager) foo() {}

// 单例实例 启动时创建
var instance *eager = &eager{data: "some data"}

// GetEagerSingleton 如果未抽象出 interface 返回值直接为 eager 则会报
// Exported function with the unexported return type
func GetEagerSingleton() Singleton {
	return instance
}
