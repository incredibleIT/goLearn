package main

import (
	"fmt"
)

// Channel是一种类型，用于在不同goroutine之间进行通信。它是一个线程安全的队列，遵循先入先出（FIFO）的原则
// 其本身带有一把锁, 多个生产者时协程安全
// 用于协程间数据传输
// 默认发送和接收操作是阻塞的，直到发送方和接收方都就绪
type Stu struct {
	Name string
	age  int
}

func main() {

	// 定义
	//var c chan int
	//c = make(chan int)

	// 使用
	c := make(chan int, 5)

	c <- 4
	c <- 1
	c <- 2
	c <- 3

	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	// 任意类型
	s := make(chan interface{}, 5)

	s <- 100
	s <- Stu{Name: "yyy", age: 30}

	fmt.Println(<-s)
	// 此处需断言<-s的类型, 之后做具体操作
	s2 := (<-s).(Stu)
	fmt.Println(s2.Name)

	// close() 关闭管道, 管道关闭后只能读不能写
	// 在循环取一个管道值之前要把管道先关闭

	close(c)

	for i := range c {
		fmt.Println(i)
	}

}
