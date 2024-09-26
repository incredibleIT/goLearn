package main

import "fmt"

// 01 一个接口可以被多个结构体所实现

type Person interface {
	run()
}

type Student struct {
	name string
}

type Teacher struct {
	name string
}

// 实现
func (stu *Student) run() {
	fmt.Printf("student %s run\n", stu.name)
}

func (teacher *Teacher) run() {
	fmt.Printf("teacher %s run\n", teacher.name)
}

// 02 一个结构体可以实现多个接口
type Animal interface {
	bark()
}

type A interface {
	run()
}

type Dog struct {
}

func (dog *Dog) bark() {
	fmt.Println("dog bark")
}

func (dog *Dog) run() {
	fmt.Println("dog run")
}

func main() {

	// 01
	var person Person

	person = &Student{name: "杨阳"}
	person.run()

	person = &Teacher{name: "YYY"}
	person.run()

	// 02

	fmt.Println("-----------------------")

	var b Animal
	b = &Dog{}
	b.bark()

	var c A
	c = &Dog{}
	c.run()

}
