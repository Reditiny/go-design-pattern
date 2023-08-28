package memento

// Caretaker 负责管理备忘录
type Caretaker struct {
	Mementos *Memento // 可以设为数组管理多个状态
}
