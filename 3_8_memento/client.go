package main

import (
	"design_pattern/3_8_memento/memento"
	"fmt"
)

func main() {
	// 创建角色
	role := memento.Role{Vit: 100, Atk: 10, Def: 10}
	// 备忘录管理器
	caretaker := memento.Caretaker{}
	// 存档
	caretaker.Mementos = role.CreateMemento()
	// 游戏失败
	role.Vit = 0
	// 读档
	role.RestoreMemento(caretaker.Mementos)
	fmt.Println(role)
}
