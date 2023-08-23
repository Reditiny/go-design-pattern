package main

import "design_pattern/2_4_facade/facade"

func main() {
	light := facade.Light{Status: false}
	airCondition := facade.AirCondition{Status: false, Temperature: 26}
	appliancesFacade := facade.AppliancesFacade{Light: light, AirCondition: airCondition}
	// 通过外观接口来隐藏内部细节
	appliancesFacade.On() // 同时打开电灯和空调
}
