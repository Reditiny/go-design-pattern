package adapter

import "fmt"

// Computer 只能读取 SD 卡
type Computer struct{}

func (c *Computer) ReadComputerSD(sdCard SDCard) {
	fmt.Println(sdCard.ReadSD())
}

func (c *Computer) WriteComputerSD(sdCard SDCard, content string) {
	sdCard.WriteSD(content)
}
