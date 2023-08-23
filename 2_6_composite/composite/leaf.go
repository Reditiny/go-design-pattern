package composite

import "fmt"

// MenuLeaf 叶子节点
type MenuLeaf struct {
	Name  string
	Level int
}

func (l *MenuLeaf) Add(menuComponent *MenuComponent) {
	fmt.Println("叶子无法添加子节点")
}

func (l *MenuLeaf) Remove(menuComponent *MenuComponent) {
	fmt.Println("叶子无法移除子节点")
}

func (l *MenuLeaf) Print() {
	for i := 0; i < l.Level; i++ {
		fmt.Print("==")
	}
	fmt.Println(l.Name)
}
