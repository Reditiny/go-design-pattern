package static

import "fmt"

// TrainStation 车站 代理对象所代表的真实对象
type TrainStation struct{}

func (s *TrainStation) Sell() {
	fmt.Println("火车站售票")
}
