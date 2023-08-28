package mediator

type Colleague interface {
	SendMessage(message string)
	RecMessage(message string)
}
