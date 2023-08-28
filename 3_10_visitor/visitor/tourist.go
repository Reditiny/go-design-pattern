package visitor

import "fmt"

// Tourist 具体访问者 实现了访问者接口
type Tourist struct {
	Name string
}

func (t *Tourist) VisitMuseum(museum *Museum) {
	fmt.Printf("游客%s参观博物馆\n", t.Name)
}

func (t *Tourist) VisitZoo(zoo *Zoo) {
	fmt.Printf("游客%s参观动物园\n", t.Name)
}
