package memento

// Role 角色原发器，用于创建和恢复备忘录
type Role struct {
	Vit int // 生命属性
	Atk int // 攻击属性
	Def int // 攻击属性
}

// CreateMemento 创建角色状态备忘录
func (r *Role) CreateMemento() *Memento {
	return &Memento{Vit: r.Vit, Atk: r.Atk, Def: r.Def}
}

// RestoreMemento 从备忘录中恢复状态
func (r *Role) RestoreMemento(m *Memento) {
	r.Vit = m.Vit
	r.Atk = m.Atk
	r.Def = m.Def
}
