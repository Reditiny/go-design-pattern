package main

import "design_pattern/2_6_composite/composite"

func main() {
	// 构建菜单树
	var c1, c2, l1, l2, l3 composite.MenuComponent
	c1 = &composite.MenuComposite{Name: "首页", Level: 0}
	c2 = &composite.MenuComposite{Name: "菜单管理", Level: 1}
	l1 = &composite.MenuLeaf{Name: "退出", Level: 1}
	l2 = &composite.MenuLeaf{Name: "角色", Level: 2}
	l3 = &composite.MenuLeaf{Name: "日志", Level: 2}
	c2.Add(&l2)
	c2.Add(&l3)
	c1.Add(&c2)
	c1.Add(&l1)
	// 打印菜单树
	c1.Print()
}
