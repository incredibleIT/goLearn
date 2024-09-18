package main

import "fmt"

func main() {

	// 值类型: int, string, bool, float []type
	x := 10
	y := x

	fmt.Printf("x的值%d, x的地址%p\n", x, &x)
	fmt.Printf("y的值%d, y的地址%p", y, &y)

	a := 100
	fmt.Printf("a的地址%p", &a)
	a = 1000
	fmt.Printf("a的地址%p", &a)

	// 引用类型, 变量存储的是地址, 这个地址存储值
	var s []int
	//s[0] = 1 切片是一个引用类型, 声明时并未开辟空间没有0索引, 此时是一个nil对象
	fmt.Println(s)

}
