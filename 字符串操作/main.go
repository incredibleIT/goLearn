package main

import "fmt"

func main() {
	s := "hello world"

	fmt.Println(s)

	fmt.Println(string(s[0]))

	// 顾头不顾尾
	fmt.Println(s[0:len(s)])
	//缺省
	fmt.Println(s[:len(s)])
	fmt.Println(s[0:])

}
