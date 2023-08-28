package mediator

import "fmt"

// M 具体中介者 实现了中介者接口 聚合了各同事对象 负责协调各个同事对象的交互
type M struct {
	HouseOwner *HouseOwner
	Tenant     *Tenant
}

func (m *M) Contact(message string, college Colleague) {
	if college == m.HouseOwner {
		m.Tenant.RecMessage(message)
	} else if college == m.Tenant {
		m.HouseOwner.RecMessage(message)
	} else {
		fmt.Printf("message:%s  发送失败\n", message)
	}
}
