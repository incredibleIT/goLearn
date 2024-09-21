package main

import "fmt"

func main() {
	/*
			匿名函数形式:
		`	func(函数参数) {
				函数体
			}

	*/

	// 匿名函数声明, 调用
	(func(x, y int) {
		fmt.Println(x + y)
	})(1, 2)

	// 匿名函数的赋值
	fun := func(x, y int) {
		fmt.Println(x + y)
	}

	fun(1, 2)

}
