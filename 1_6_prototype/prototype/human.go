package prototype

import "time"

type Human struct {
	Flesh string
	Soul  []time.Time
}

func (h *Human) Clone() Cloneable {
	newOne := Human{}
	newOne.Flesh = h.Flesh
	// 切片类型进行深拷贝
	newOne.Soul = make([]time.Time, len(h.Soul))
	copy(newOne.Soul, h.Soul)
	return &newOne
}
