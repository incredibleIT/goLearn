package main

import (
	"fmt"
	"strconv"
	"time"
)

// 测试的调用函数
func f() {
	fmt.Println("f被调用")
	time.Sleep(time.Second * 3)
}

/*

	我对这段函数是这样理解的,
	返回的匿名函数是一个闭包函数他维护了一个自由变量count
	在接收counter返回的闭包函数时x并不会因为counter的调用结束而被释放,
	当闭包函数调用时x依然在栈中, 并受闭包函数影响, 每一次调用counter都会返回一个闭包函数他们拥有不同的底层栈, 所以维护了不互相影响的x
*/
// counter可以叫做装饰函数
func counter(f func()) func() int {
	count := 0

	return func() int {
		f()
		count++
		fmt.Println(count)

		return count

	}
}

// 一个计算函数运行时间的装饰函数
func timer(fn func()) func() int64 {

	return func() int64 {

		start := time.Now().Unix()
		fn()
		end := time.Now().Unix()
		return end - start

	}

}

func main() {
	f2 := counter(f)
	f2()
	f2()
	c1 := f2()
	fmt.Println("c1 : " + strconv.Itoa(c1))

	fmt.Println("==========")

	f3 := counter(f)
	f3()
	f3()
	c2 := f3()
	fmt.Println("c2 : " + strconv.Itoa(c2))

	f4 := timer(f)
	fmt.Println(f4())
}
