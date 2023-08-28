package observer

import "fmt"

// Observer 观察者对象 也可以抽象出接口分为抽象观察者与具体观察者
type Observer struct {
	Name string
}

func (o *Observer) Update(message string) {
	fmt.Printf("%s接收到新消息:%s\n", o.Name, message)
}
