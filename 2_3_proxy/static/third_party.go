package static

import "fmt"

// ThirdParty 第三方售票代理 可以访问、控制或扩展真实主题（车站）的功能
type ThirdParty struct {
	TrainStation
}

func (c *ThirdParty) Sell() {
	fmt.Println("第三方收取服务费")
	// 真实业务 代理可在其前后添加额外业务逻辑
	c.TrainStation.Sell()
	fmt.Println("第三方提供额外服务")
}
