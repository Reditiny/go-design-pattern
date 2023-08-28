package iterator

// StudentAggregator 抽象聚合接口 定义存储、添加、删除聚合元素以及创建迭代器对象的接口
type StudentAggregator interface {
	AddStudent(stu *Student)
	RemoveStudent(stu *Student)
	GetStudentIterator() StudentIterator
}
