package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明map var mapName[key_type] value_type

	// 当声明未赋值时, 不会开辟空间
	var mapName map[string]string
	fmt.Println(mapName, reflect.TypeOf(mapName))

	// make函数开辟, 第二个参数是map的cap, 当map增长到容量上限时, 如果再增加新的键值对, map的大小会自动加一, 即使知道大概容量也最好先标明
	mapName = make(map[string]string, 100)
	fmt.Println(mapName, reflect.TypeOf(mapName))
	mapName["name"] = "hello"
	mapName["age"] = "world"

	// map 类型具有无序性
	fmt.Println(mapName)

	// 直接声明赋值
	mapName2 := map[string]string{"name": "yy", "age": "20", "gender": "nan"}
	fmt.Println(mapName2)

	// map 访问
	fmt.Println(mapName2["gender"])
	fmt.Println(mapName2["name"])
	fmt.Println(mapName2["age"])

	// range循环
	for k, v := range mapName2 {
		fmt.Println(k, v)
	}

	// 删除元素delete
	delete(mapName2, "name")
	fmt.Println(mapName2)

}
