package main

import "design_pattern/3_10_visitor/visitor"

func main() {
	// 具体元素 博物馆和动物园
	museum := &visitor.Museum{Name: "自然博物馆", Fee: 0}
	zoo := &visitor.Zoo{Name: "野生动物园", Fee: 100}
	// 游客 可以定义其他类型人员进行不同方式的访问
	tourist := visitor.Tourist{Name: "red"}
	// 访问
	tourist.VisitZoo(zoo)
	tourist.VisitMuseum(museum)
}
