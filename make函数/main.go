package main

import "fmt"

func main() {
	// 针对引用数据类型, 需要声明开辟空间的方法
	s := make([]int, 5) // int切片类型, 长度和容量都是5
	s[0] = 100
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	a := make([]int, 5, 10) // 长度是5, 容量是10, 可进行扩容

	fmt.Println(len(a), cap(a))

}
