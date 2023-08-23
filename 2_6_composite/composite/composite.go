package composite

import "fmt"

// MenuComposite 组合节点 其中可以包括叶子节点和组合节点
type MenuComposite struct {
	Children []*MenuComponent
	Name     string
	Level    int
}

func (l *MenuComposite) Add(menuComponent *MenuComponent) {
	// 这里省略了判断 level 是否正确的逻辑
	l.Children = append(l.Children, menuComponent)
	fmt.Println("添加成功")
}

func (l *MenuComposite) Remove(menuComponent *MenuComponent) {
	// 省略删除逻辑
	fmt.Println("删除成功")
}

func (l *MenuComposite) Print() {
	for i := 0; i < l.Level; i++ {
		fmt.Print("==")
	}
	fmt.Println(l.Name)
	for _, c := range l.Children {
		(*c).Print()
	}
}
