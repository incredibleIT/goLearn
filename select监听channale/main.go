package main

import (
	"fmt"
	"time"
)

func f1(c chan<- int) {

	time.Sleep(2 * time.Second)

	c <- 1

}

func f2(c chan<- int) {
	time.Sleep(3 * time.Second)

	c <- 2
}

func main() {

	c1 := make(chan int, 100)
	//
	//c1 <- 1
	//c1 <- 2
	//c1 <- 3
	//
	c2 := make(chan int, 100)
	//c2 <- 4
	//c2 <- 5
	//c2 <- 6

	// 注意当多个case都有数据的话, 会随机取一个执行
	//select {
	//case v1 := <-c1: // 监听c1读
	//	fmt.Println("<- c1 :", v1)
	//case v2 := <-c2: // 监听c2读
	//	fmt.Println("<- c2 :", v2)
	//default:
	//	fmt.Println("default")
	//}

	go f1(c1)
	go f2(c2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)

		}
	}

}
