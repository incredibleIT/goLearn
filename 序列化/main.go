package main

import (
	"encoding/json"
	"fmt"
)

type Addr struct {
	City string
}

// 首先结构体的定义
type Student struct {
	Number int
	Name   string
	Age    int
	Addr   Addr
}

func main() {
	/*
		序列化: 通过某种方式把数据结构或对象写入到磁盘文件中或通过网络传到其他节点的过程
		反序列化: 把磁盘中的对象或者网络节点中传输的数据恢复为数据对象的过程
	*/

	// 序列化
	map1 := map[int]string{1: "hello", 2: "world"}

	map1Json, _ := json.Marshal(map1)

	fmt.Println(string(map1Json))

	student1 := Student{2, "yangyang", 19, Addr{City: "beijing"}}

	studentJson, _ := json.Marshal(student1)

	fmt.Println(string(studentJson))

	// 反序列化

	var map2 map[int]string

	err := json.Unmarshal(map1Json, &map2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(map2)

}
