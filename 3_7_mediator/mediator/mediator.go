package mediator

type Mediator interface {
	Contact(message string, college Colleague)
}
