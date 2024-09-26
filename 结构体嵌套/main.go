package main

import "fmt"

/*
	并不是一个结构体实现一个接口的所有方法才叫做这个结构体实现了这个接口,
	一个结构体继承自另一个实现了这个接口一部分方法的结构体, 而本结构体实现了剩余的方法, 这样也可以叫做实现了这个接口
*/

type Chuifengji interface {
	chui()

	dry()
}

type Dianji struct {
}

type Chuifeng struct {
	*Dianji
}

func (d *Dianji) dry() {
	fmt.Println("烘干")
}

func (c *Chuifeng) chui() {
	fmt.Println("吹风")
}

func main() {

	var c Chuifengji

	// 相当于实现了Chuifengji接口
	c = &Chuifeng{}

	c.dry()

	c.chui()

}
