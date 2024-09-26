package main

import "fmt"

type Animal interface {
	sleep02()
}

func (d Dog) sleep() {
	fmt.Println("dog sleep")

}

func (d *Dog) sleep02() {
	fmt.Println("dog sleep")

}

type Dog struct {
}

func main() {

	var animal Animal

	var dog = Dog{}

	// 此时dog可以赋值给animal这个接口类型
	//animal = dog

	// 发现结构体的指针类型也可以赋值给其接口类型, 因为语法糖的原因如果传入的是地址GO会自动作取值操作
	//animal = &dog

	//但如果方法接收器要求调用的类型是指针类型则必须传入指针类型, 因为go不会作值类型转为指针类型的操作
	//animal = dog (错误的)

	animal = &dog

	animal.sleep02()

}
