package mediator

import "fmt"

type Tenant struct {
	Name     string
	Mediator Mediator
}

func (t *Tenant) SendMessage(message string) {
	t.Mediator.Contact(message, t)
}
func (t *Tenant) RecMessage(message string) {
	fmt.Printf("%v收到消息:%s\n", t, message)
}
