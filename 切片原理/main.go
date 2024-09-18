package main

import "fmt"

func main() {

	// 切片不额外开辟空间, 其具有三个部分, 1.起始指针2.长度3.容量
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr[2:6]

	fmt.Println(s1, len(s1), cap(s1)) // 长度是切片起始位置到切片结尾长度, 容量是切片起始到所切数组结尾长度cap()容量

}
