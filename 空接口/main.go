package main

import "fmt"

// 空接口中没有定义方法, 所以可以认为任何一种类型都是空接口的一个实现
// 换言之, 空接口类型可以接受任意一种类型的变量
type empty interface {
}

type Person struct {
	name string
}

func main() {

	var em empty

	em = 1
	fmt.Println(em)

	em = Person{"yang"}

	fmt.Println(em)

	// 当将基本类型数据传给空接口变量时, 变量会保存一个副本在接口中
	// 但对于带有结构的类型, 接口则不会存储其结构, 仅仅只是存值, 这时必须使用类型断言来获取原始类型，才能访问结构体的字段。
	if val, ok := em.(Person); ok {
		fmt.Println(val.name)
	} else {
		fmt.Println("err")
	}

	// 在switch语句中可以使用空接口变量.(type)来快速判断类型
	switch em.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Person:
		fmt.Println("Person")

	}

}
