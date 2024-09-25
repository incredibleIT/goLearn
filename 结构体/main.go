package main

import (
	"fmt"
	"reflect"
)

// 声明结构体: (类的概念) 值类型
type Student struct {
	name   string
	age    int
	gender byte
	score  []float64
}

func main() {

	// 实例化1
	var stu Student   // 声明
	stu.name = "yang" // 初始化
	fmt.Println(stu)

	// 实例化2
	s := Student{name: "yang", age: 12, gender: 1, score: []float64{1.0, 2.0, 3.0}}

	fmt.Println(s, reflect.TypeOf(s))

	s2 := Student{"yang", 11, 2, []float64{1, 2, 3}}

	fmt.Println(s2)

	// 实例化3 &结构体{}
	var a *Student

	a = &Student{name: "yang"}

	// 暗箱操作, 在go中访问结构体指针的成员变量时可以使用., 这是因为为了方便开发者访问结构体指针的成员变量可以像访问结构体的成员变量一样简单, 将instance.Name形式转换为(*instance).Name
	fmt.Println((*a).name, reflect.TypeOf(a)) // 正常写法

	fmt.Println(a.age) // 语法糖 a直接.就可以取到结构体中的属性

	b := &Student{name: "yang", age: 12, gender: 1, score: []float64{1.0, 2.0, 3.0}}

	change02(b)

	fmt.Println("b: ", *b)

	// 实例化4  new

	stu1 := new(Student) // 等同于 &Student{}

	fmt.Println(*stu1, reflect.TypeOf(stu1)) // &{ 0 0 []} *main.Student

}

func change02(b *Student) {
	b.age = 1
}

func change(b Student) {
	b.age = 1
}
