package adapter

// TFCard 原接口
type TFCard interface {
	ReadTF() string
	WriteTF(content string)
}
