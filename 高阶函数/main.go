package main

import (
	"fmt"
	"strconv"
	"time"
)

func a() {
	fmt.Println("hello a start")
	time.Sleep(2 * time.Second)
	fmt.Println("hello a end")
}

func b() {
	fmt.Println("hello b start")
	time.Sleep(3 * time.Second)
	fmt.Println("hello b end")
}

// 测试时间的高阶函数
// 以函数作参数
func timer(f func()) string {
	start := time.Now().Unix()
	f()
	end := time.Now().Unix()
	return "消耗时间: " + strconv.Itoa(int(end-start))
}

// 一函数作为返回值
func outer() func() {
	// 声明匿名函数, 作为参返回
	return func() {
		fmt.Println("inner outer")
	}
}

func main() {

	/*
		高阶函数满足以下条件之一:
							1. 以函数作为参数
							2. 以函数作为返回值
	*/

	// 本质上就是函数作为参数传递
	f := func() { fmt.Println("hello world") }

	fmt.Println(timer(f))
	fmt.Println(timer(a))
	fmt.Println(timer(b))

	// 直接传匿名函数
	fmt.Println(timer(func() { fmt.Println("hello world") }))

	// 作返回值返回
	outer()()
}
