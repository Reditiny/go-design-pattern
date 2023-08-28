package mediator

import "fmt"

type HouseOwner struct {
	Name     string
	Mediator Mediator
}

func (h *HouseOwner) SendMessage(message string) {
	h.Mediator.Contact(message, h)
}
func (h *HouseOwner) RecMessage(message string) {
	fmt.Printf("%v收到消息:%s\n", h, message)
}
