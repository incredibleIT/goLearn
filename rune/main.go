package main

import "fmt"

// rune用来处理长度大于1字节, 不超过四字节的字符类型
// rune通过unicode码点来对字符进行存储, 可以表示世界上任意一个字符
func main() {

	// 记录长度
	var s string = "杨阳"
	fmt.Println(len(s)) // 6

	s2 := []rune("杨阳")
	fmt.Println(len(s2))

	// 截取
	fmt.Println(s[:2])

	fmt.Println(string(s2[:2]))

}
