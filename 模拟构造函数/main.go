package main

import "fmt"

type Student struct {
	name   string
	age    int
	gender byte
	score  [3]int
}

/*
	这两个函数的主要区别在于返回值的类型和内存管理。

	- **`NewStudent01`** 返回一个指向 `Student` 结构体的指针。这意味着在函数外部对 `Student` 的修改会影响到原始对象，且可以节省内存（适合大的结构体）。

	- **`NewStudent02`** 返回 `Student` 的值，这意味着函数外部获得的是一个副本。修改这个副本不会影响原始对象，且每次调用都会创建一个新的结构体实例。

	使用哪个函数取决于你是否需要共享对象状态或希望确保对象独立。

*/
func NewStudent01(name string, age int, gender byte, score [3]int) *Student {
	return &Student{
		name:   name,
		age:    age,
		gender: gender,
		score:  score,
	}
}

func NewStudent02(name string, age int, gender byte, score [3]int) Student {
	return Student{
		name:   name,
		age:    age,
		gender: gender,
		score:  score,
	}
}

func NStudent() *Student {
	return new(Student)
}

func main() {

	s := NewStudent02("yang", 1, 1, [3]int{1, 2, 3})

	fmt.Println(s)

	s.name = "sdfgdrfg"

	fmt.Println(s)

}
