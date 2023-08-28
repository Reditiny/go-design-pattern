package responsibility_chan

import "fmt"

// Manager 具体处理者2 可以同意请求或将请求传递给下一个处理者
type Manager struct {
	Max         int
	NextHandler Handler
}

func (m *Manager) Handle(days int) {
	if days <= m.Max {
		fmt.Println("Manager approve request")
	} else {
		fmt.Println("Manager need to send request to Boss")
		m.NextHandler.Handle(days)
	}
}
