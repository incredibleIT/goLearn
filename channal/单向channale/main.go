package main

import "fmt"

// 对于单向channale可由参数处指定在一个函数中的读写级别
func read(c <-chan int) {
	fmt.Println(<-c)
}

func write(c chan<- int) {
	c <- 1

}

func main() {

	// 声明
	c1 := make(<-chan int, 3) // 只读channale
	fmt.Println(<-c1)

	c2 := make(chan<- int, 3) // 只写channale
	c2 <- 2

	c := make(chan int, 3) // 双向channale

	// 定义读写级别
	read(c)
	write(c)

}
