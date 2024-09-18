package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	//strings.Index()
	//strings.LastIndex()
	//strings.HasPrefix()
	//strings.HasSuffix()
	//strings.Contains()
	//strings.Split()
	//strings.Join()

	s := "hello,world"

	// 长度
	fmt.Println(len(s))
	// 分割
	split := strings.Split(s, ",")
	fmt.Println(split)
	// strings.Join(): 将数组拼接成字符串
	join := strings.Join(split, " ")
	fmt.Println(join, reflect.TypeOf(join))
	// 是否包含
	fmt.Println(strings.Contains(s, "hellw"))

	// 子串索引, 未找到返回-1, 第一次出现位置以及最后一次出现位置
	index := strings.Index(s, "hello")
	fmt.Println(index)

	lastIndex := strings.LastIndex(s, "l")
	fmt.Println(lastIndex)

	// 是否以什么开头或结尾
	prefix := strings.HasPrefix(s, "0")
	fmt.Println(prefix)

	suffix := strings.HasSuffix(s, "l")
	fmt.Println(suffix)
}
