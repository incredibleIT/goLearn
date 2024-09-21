package main

import "fmt"

// 值类型: 基本数据类型, 数组类型, 结构体

// 引用类型: 切片, map, 指针类型

// a 是局部变量
func f(a int) {
	a = 100
}

// 引用传递, 传递的是地址
func f2(s []int) {
	fmt.Printf("%p\n", s)
	s[0] = 100
	s = append(s, 1)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
}

func main() {

	// 1  值类型传递, 产生的是拷贝传递
	x := 10
	y := x
	x = 100
	fmt.Println(x, y)

	// 2
	a := 1
	f(a) // 值传递
	fmt.Println(a)

	// 3 将s传递,传递的是存s的地址, 两个函数共用一套切片, 不过若切片容量不足在函数中append的时候会产生新的切片产生新的地址, 不过这不会影响原始切片
	s := make([]int, 3, 5)
	fmt.Printf("%p\n", s)
	f2(s)
	fmt.Println(s)
}
