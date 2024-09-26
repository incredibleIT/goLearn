package main

import "fmt"

type Run interface {
	run()
}

type Jump interface {
	jump()
}

// 这样写很麻烦
//type Animal interface {
//	run()
//	jump()
//}

// 实现Run Jump两个接口就间接实现了Animal接口
type Animal interface {
	Run
	Jump
}

type Dog struct {
}

func (d *Dog) run() {
	fmt.Println("running.....")
}

func (d *Dog) jump() {
	fmt.Println("jumping.....")
}

func main() {

	var v *Dog
	var a Animal

	a = v

	a.jump()
	a.run()

}
