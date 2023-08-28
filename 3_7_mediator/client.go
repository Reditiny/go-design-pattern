package main

import "design_pattern/3_7_mediator/mediator"

func main() {
	// 中介

	m := &mediator.M{}
	// 房主
	owner := &mediator.HouseOwner{Name: "John"}
	// 租客
	tenant := &mediator.Tenant{Name: "red"}
	// 中介管理双方通信
	m.HouseOwner = owner
	m.Tenant = tenant
	owner.Mediator = m
	tenant.Mediator = m
	// 双方通信
	tenant.SendMessage("多少钱一个月")
	owner.SendMessage("2000一个月")

}
