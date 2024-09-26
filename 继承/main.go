package main

import "fmt"

// 结构体继承: 通过匿名成员变量来实现的
type Person struct {
	Name string
}

// Student继承了Person, 即Person作为Student的子类
type Student struct {
	sex byte
	*Person
}

type Teacher struct {
	*Student
	*Person
}

func (p *Person) run() {
	fmt.Println(p.Name, "running ......")
}

func (s *Student) learn() {
	fmt.Println(s.Name, "learn ......")
}

func NewStudent(sex byte, p *Person) *Student {
	return &Student{
		sex: sex,

		Person: p,
	}
}

func main() {

	student := NewStudent('1', &Person{Name: "yy"})
	student.learn()

}
