package adapter

// SDCard 目标接口
type SDCard interface {
	ReadSD() string
	WriteSD(content string)
}
