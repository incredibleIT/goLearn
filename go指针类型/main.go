package main

import "fmt"

func main() {

	// x是一个整型变量
	var x = 10

	// 取址: &   取值: *
	fmt.Println(x)

	//指针类型, 存储的是地址
	var p *int = &x
	fmt.Println(*p)

	s := "hello world"

	var a *string = &s
	fmt.Println((*a)[1:])
}
