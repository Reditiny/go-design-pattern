package facade

type AppliancesFacade struct {
	Light
	AirCondition
}

func (a *AppliancesFacade) On() {
	a.Light.OpenLight()
	a.AirCondition.OpenAirCondition()
}

func (a *AppliancesFacade) Off() {
	a.Light.OffLight()
	a.AirCondition.OffAirCondition()
}
