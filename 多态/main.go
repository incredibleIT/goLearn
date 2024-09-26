package main

import "fmt"

// go的多态是由interface实现的, 一个结构体定义了与interface中所有方法同名的方法接收器即认为这个结构体实现了这个接口
// 并可由接口类型接收此结构体的变量(对象), 这就是所谓的多态
// 接口类型的变量调用方法时显现出不同的状态

// 多态
type Pay interface {
	pay() string
}

type WeiXinPay struct {
}

type YinLianPay struct {
}

type ALiPay struct {
}

func (w WeiXinPay) pay() string {
	fmt.Println("微信支付")
	return "微信"
}

func (y YinLianPay) pay() string {
	fmt.Println("银联支付")
	return "银联"
}

func (a ALiPay) pay() string {
	fmt.Println("阿里支付")
	return "阿里"
}

func Payer(p Pay) string {

	return p.pay()

}

func main() {

	p01 := WeiXinPay{}

	fmt.Println(p01.pay())

	p02 := YinLianPay{}
	fmt.Println(p02.pay())

	p03 := ALiPay{}
	fmt.Println(p03.pay())

}
