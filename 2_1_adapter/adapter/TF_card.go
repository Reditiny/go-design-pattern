package adapter

import "fmt"

// SpecificTF 原接口具体实现
type SpecificTF struct{}

func (s *SpecificTF) ReadTF() string {
	return "reading TF card"
}

func (s *SpecificTF) WriteTF(content string) {
	fmt.Println("write some content into TF card")
}
