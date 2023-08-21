package singleton

import "sync"

// 懒汉式单例 包私有
type lazy struct {
	data string
}

func (e *lazy) foo() {}

// 单例实例 都一次使用时创建
var (
	lazyInstance *lazy
	once         sync.Once // 实现线程安全的单次执行
)

// GetLazySingleton 如果未抽象出 interface 返回值直接为 lazy 则会报
// Exported function with the unexported return type
func GetLazySingleton() Singleton {
	once.Do(func() { // Do 仅在第一次调用时执行
		lazyInstance = &lazy{data: "some data"}
	})
	return instance
}
