package main

import (
	"fmt"
	"reflect"
)

func main() {

	var arr1 [3]int

	// 默认为零值
	fmt.Println(arr1)

	// 通过索引赋值(初始化)
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)

	// 声明初始化
	var arr2 = [3]int{1, 2, 3}
	var arr3 = [...]int{1, 2, 3}
	fmt.Println(arr2, " ", arr3)

	// 索引声明
	var arr4 = [3]int{0: 1, 2: 100}
	fmt.Println(arr4)

	arr5 := [4]int{1, 2, 3, 4}
	fmt.Println(arr5)

	// 访问数据, 索引,切片, 循环

	arr6 := [3]string{"hello", "world", "yyy"}
	// 索引
	fmt.Println(arr6[1])

	// 切片
	fmt.Println(arr6[0:2], "typeofarr6[0 : 2]:", reflect.TypeOf(arr6[0:2])) // arr6[0:2] 是切片类型

	// 循环
	for i := 0; i < len(arr6); i++ {
		fmt.Println(arr6[i])
	}

	// 遍历循环
	// 一个变量拿的是索引
	for i := range arr6 {

		fmt.Println(i, arr6[i])
	}

	//两个变量拿的是索引和索引对应的值
	for i, v := range arr6 {
		fmt.Println(i, v)
	}
}
