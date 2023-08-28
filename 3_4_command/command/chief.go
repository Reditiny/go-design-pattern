package command

import "fmt"

// Chief 接收者
type Chief struct {
}

func (c *Chief) MakeFood(name string, count int) {
	fmt.Printf("厨师制作了%d份%s\n", count, name)
}
