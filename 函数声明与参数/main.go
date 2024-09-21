package main

import (
	"fmt"
	"reflect"
)

// 函数定义
func fun1() {

	fmt.Println("hello world")

}

func fun2(a int) {

	fmt.Println(a)
}

func fun3(a, b int) {
	fmt.Println(a, b)
}

func fun4(a int, b int) {
	fmt.Println(a, b)
}

func fun5(x ...int) { // x是切片
	for i, j := range x {
		fmt.Println(i, j)
	}

	fmt.Println(reflect.TypeOf(x))
}

func main() {
	fun1()
	fun2(1)
	fun3(1, 2)
	fun5(1, 2, 3, 4, 5, 6)
}
