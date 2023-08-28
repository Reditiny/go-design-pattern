package responsibility_chan

import "fmt"

// GroupLeader 具体处理者1 可以同意请求或将请求传递给下一个处理者
type GroupLeader struct {
	Max         int
	NextHandler Handler
}

func (g *GroupLeader) Handle(days int) {
	if days <= g.Max {
		fmt.Println("Group leader approve request")
	} else {
		fmt.Println("Group leader need to send request to Manager")
		g.NextHandler.Handle(days)
	}
}
