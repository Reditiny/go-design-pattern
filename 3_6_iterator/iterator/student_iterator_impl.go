package iterator

// StudentIteratorImpl 具体迭代器 完成对聚合对象的遍历并记录遍历的当前位置。
type StudentIteratorImpl struct {
	studentList []*Student
	position    int
}

func (i *StudentIteratorImpl) Next() *Student {
	cur := i.position
	i.position += 1
	return i.studentList[cur]
}

func (i *StudentIteratorImpl) HasNext() bool {
	return i.position < len(i.studentList)
}
