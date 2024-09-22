package main

import "fmt"

func main() {

	// 1延迟到函数结束运行
	//defer fmt.Println("hello world1")
	//fmt.Println("hello world2")
	//fmt.Println("************************************")

	// 2多个defer的执行顺序: 先注册的后执行
	defer fmt.Println("hello world3")
	defer fmt.Println("hello world4")
	defer fmt.Println("hello world5")
	fmt.Println("hello world6")

	// defer拷贝机制, 注册时注册的不是函数名, 而是当时的函数体
	f := func() {
		fmt.Println("f1")
	}

	defer f()

	f = func() {
		fmt.Println("f2")
	}

	// 3 将函数体整体保存起来, x = 10保存起来
	x := 10
	defer func(x int) {
		fmt.Println(x)
	}(x)

	x++

	// 4 引用的一套变量
	y := 10
	defer func() {
		fmt.Println(y)
	}()
	y++

}
