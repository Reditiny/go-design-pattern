package main

import (
	"design_pattern/3_6_iterator/iterator"
	"fmt"
)

func main() {
	student1 := &iterator.Student{Name: "red", Id: 1}
	student2 := &iterator.Student{Name: "min", Id: 1}
	student3 := &iterator.Student{Name: "lee", Id: 1}
	// 学生集合
	studentList := iterator.StudentAggregatorImpl{}
	studentList.AddStudent(student1)
	studentList.AddStudent(student2)
	studentList.AddStudent(student3)
	// 获取迭代器 利用迭代器变量集合
	studentIterator := studentList.GetStudentIterator()
	for studentIterator.HasNext() {
		fmt.Println(studentIterator.Next())
	}

}
