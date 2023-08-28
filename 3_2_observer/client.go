package main

import "design_pattern/3_2_observer/observer"

func main() {
	// 发布方
	s := observer.Subscription{}
	// 新增订阅方
	o1 := &observer.Observer{Name: "red"}
	o2 := &observer.Observer{Name: "lee"}
	o3 := &observer.Observer{Name: "joe"}
	s.Attach(o1)
	s.Attach(o2)
	s.Attach(o3)
	// 发布消息1
	s.Notify("welcome here")
	// 取消订阅
	s.Detach(o2)
	s.Detach(&observer.Observer{Name: "ming"}) // 一个未订阅的用户取消不会有任何影响（取决于内部实现）
	// 发布消息2
	s.Notify("here is the first message")

}
