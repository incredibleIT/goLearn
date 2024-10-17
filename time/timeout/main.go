package main

import (
	"fmt"
	"time"
)

var t = make(chan int)

// 处理网络请求数据库查询时, 超时控制非常重要
func main() {

	go func() {
		time.Sleep(3 * time.Second)
		t <- 100
	}()

	select {
	case res := <-t:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}

	// time.Sleep()的替代方法. time.After()可以实现类似Sleep的效果, 它通过通道让其他任务同步等待

	<-time.After(time.Second * 5)
	fmt.Println("hello world")

}
