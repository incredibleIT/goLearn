package main

import "fmt"

func main() {

	c1 := make(chan int, 100)

	c1 <- 1
	c1 <- 2
	c1 <- 3

	c2 := make(chan int, 100)
	c2 <- 4
	c2 <- 5
	c2 <- 6

	// 注意当多个case都有数据的话, 会随机取一个执行
	select {
	case v1 := <-c1: // 监听c1读
		fmt.Println("<- c1 :", v1)
	case v2 := <-c2: // 监听c2读
		fmt.Println("<- c2 :", v2)
	default:
		fmt.Println("default")
	}

}
