package main

import "fmt"

type Addr struct {
	City string
}

// 首先结构体的定义
type Student struct {
	Number int
	Name   string
	Age    int
	Addr
}

func main() {
	// 四种结构体初始化的方式
	// 1
	var student1 Student
	student1.Number = 1
	student1.Name = "yang"
	student1.Age = 18
	fmt.Println(student1)

	// 2
	student2 := Student{Number: 2, Name: "yangyang", Age: 12}
	// 或者
	student3 := Student{2, "yangyang", 12, Addr{}}
	fmt.Println(student2, student3)

	// 3
	student4 := &Student{Number: 2, Name: "yangyang", Age: 12}

	fmt.Println(student4.Age) // 语法糖

	// 4
	student5 := new(Student)
	fmt.Println(student5.Age) // 语法糖

	// 可以模拟构造函数的形式, 分为值拷贝和地址拷贝
	student6 := NewStudent(2, "yangyang", 12)
	fmt.Println(student6.Age)

	// 相关方法的定义
	student6.learn()

	// 匿名成员变量: 函数名默认为类型名, 在当前结构体中找不到的变量默认去匿名成员变量的struct中去找
	student6.Addr = Addr{City: "beijing"}

	fmt.Println(student6.City)

}

func NewStudent(n int, name string, age int) *Student {
	return &Student{
		Number: n,
		Name:   name,
		Age:    age,
	}
}

func (s Student) learn() {

	fmt.Println(s.Name, "学习了")

}
