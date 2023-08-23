package facade

type AirCondition struct {
	Status      bool
	Temperature int // 本例中不重要
}

func (l *AirCondition) OpenAirCondition() {
	l.Status = true
}

func (l *AirCondition) OffAirCondition() {
	l.Status = false
}
