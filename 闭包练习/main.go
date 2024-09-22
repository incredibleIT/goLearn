package main

import "fmt"

func makeFun(i int) func() {
	return func() {
		fmt.Println(i)
	}
}

func main() {
	//var fn [10]func()
	//
	//for i := 0; i < len(fn); i++ {
	//	//c := i  解决办法1, 每次形成的都是一个全新的变量,匿名函数捕获的是变量的引用
	//	fn[i] = func() {
	//		//fmt.Println(c)
	//
	//		fmt.Println(i)
	//	}
	//}
	//
	//for _, v := range fn {
	//	v()
	//}

	var fn [10]func()

	for i := 0; i < 10; i++ {
		fn[i] = makeFun(i)
	}

	for _, f := range fn {
		f()
	}

}
