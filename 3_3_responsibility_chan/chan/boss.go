package responsibility_chan

import "fmt"

// Boss 具体处理者3 最终同意请求或拒绝请求
type Boss struct {
	Max int
}

func (b *Boss) Handle(days int) {
	if days <= b.Max {
		fmt.Println("Boss approve request")
	} else {
		fmt.Println("Boss reject request")
	}
}
