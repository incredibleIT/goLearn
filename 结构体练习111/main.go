package main

import "fmt"

type rect struct {
	width, height int
}

func newRect01(width, height int) *rect {

	return &rect{
		width:  width,
		height: height,
	}

}

func newRect02(width, height int) rect {

	return rect{
		width:  width,
		height: height,
	}

}

func (r *rect) change() {

	r.height--

}

func (r *rect) area() int {
	return r.width * r.height
}

func main() {

	//r := &rect{width: 10, height: 5}

	//fmt.Println(r.area()) // 自动取地址

	//rect01 := newRect01(1, 1)

	rect02 := newRect02(1, 1)

	rect04 := rect02

	//rect03 := rect01

	rect02.change()

	fmt.Println(rect04)
}
