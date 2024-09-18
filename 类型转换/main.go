package main

import (
	"fmt"
	"strconv"
)

func main() {

	// go不存在隐式转换

	// 不允许
	//var s = "hello"

	//fmt.Println(s + 1)

	x := 100
	fmt.Println(string(x)) //转成对应字符

	fmt.Println(strconv.Itoa(x)) //转成字符串 strconv.Itoa(int)

	x1 := "100"
	r1, _ := strconv.Atoi(x1) //字符串转整型 strconv.Atoi(string) 返回一个整数, 一个错误
	fmt.Println(r1)

	f1 := "12.9"
	s1, _ := strconv.ParseFloat(f1, 64)
	fmt.Println(s1)
}
