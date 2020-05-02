package _func

import (
	"fmt"
	"testing"
)
//测试程序目的：如何使用printStudent这个函数?


type Student struct {
	Name string
	Age int
}

func printStudent(stuendtPtr interface{})  {
	stu := *(stuendtPtr.(*Student)) // 问题难点
	fmt.Println(stu.Name)
	fmt.Println(stu.Age)
}

// 解决方案
func TestTMain(t *testing.T) {
	student := Student{}
	student.Age = 18
	student.Name = "tao"

	// 关键所在
	convertFun := func(stu Student)(interface{}) {
		return &stu
	}
	stuResult := convertFun(student)

	//stu01 := *(stuResult.(*Student))
	//fmt.Println(stu01)
	printStudent(stuResult)
}
