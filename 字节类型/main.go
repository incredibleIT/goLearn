package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*
		go语言的编码与解码
		var a int8  bbbbbbbb
					00000000


		var b int16  bbbbbbbb bbbbbbbb
					 00000000 00000000
	*/

	var x byte
	var y uint8 // 在go中uint8 和 byte基本上可以划等号

	x = 'a'
	fmt.Println(x, reflect.TypeOf(x)) // 97 uint8

	y = 99
	fmt.Printf("%c\n", y)

	var s string = "hello world" // 默认UTF8编码
	fmt.Println(s, reflect.TypeOf(s))

	// 字符串(string)转换为字节串([]byte)
	b := []byte(s)
	fmt.Println(b, reflect.TypeOf(b))

	// 字节串转换字符串
	fmt.Println(string(b))

}
