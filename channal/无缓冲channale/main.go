package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int)

	// 报错 : all goroutines are asleep - deadlock!
	// 原因是, make(chan int)make的是一个无缓冲通道,无缓冲通道只有在有人接收值时才可以发送值
	//c <- 1
	//c <- 2
	//c <- 3

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			x, ok := <-c
			fmt.Println(x)
			if !ok {
				return
			}

		}
	}()

	var x = 1

	for {
		c <- x
		x++
	}

}
