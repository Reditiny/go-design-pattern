package facade

type Light struct {
	Status bool
}

func (l *Light) OpenLight() {
	l.Status = true
}

func (l *Light) OffLight() {
	l.Status = false
}
