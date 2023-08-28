package iterator

// StudentAggregatorImpl  实现抽象聚合 包含一种集合（列表） 并能返回一个具体迭代器的实例。
type StudentAggregatorImpl struct {
	studentList []*Student
}

// AddStudent 添加学生到列表末尾
func (s *StudentAggregatorImpl) AddStudent(student *Student) {
	s.studentList = append(s.studentList, student)
}

// RemoveStudent 从列表中删除指定学生 学生若本身不存在则无影响
func (s *StudentAggregatorImpl) RemoveStudent(student *Student) {
	newList := make([]*Student, len(s.studentList))
	count := 0
	for _, stu := range s.studentList {
		if stu != student {
			newList[count] = stu
			count++
		}
	}
	s.studentList = newList[0:count]
}

// GetStudentIterator 获取迭代器
func (s *StudentAggregatorImpl) GetStudentIterator() StudentIterator {
	return &StudentIteratorImpl{studentList: s.studentList, position: 0}
}
