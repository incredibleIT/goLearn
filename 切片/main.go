package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 声明数组
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr, reflect.TypeOf(arr))

	// 切片有两种创建方式
	// 切割数组或切片
	// 切数组
	s1 := arr[2:5]
	fmt.Println(s1, reflect.TypeOf(s1))

	// 从切片上切
	s2 := s1[0:2]
	fmt.Println(s2, reflect.TypeOf(s2))
	//具有引用关系
	//s1[0] = 100
	//fmt.Println(arr)
	//fmt.Println(s1, reflect.TypeOf(s1))

	// 声明

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(slice, reflect.TypeOf(slice), len(slice))

}
