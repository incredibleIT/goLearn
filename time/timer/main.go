package main

import (
	"fmt"
	"time"
)

/*
一旦定时器触发了, time.C会接收到一个值, 如果试图在之后调用Stop(), 会false因为定时器已经触发了
*/
func main() {
	// 1. go的time.Timer 可用来设置一次性的定时操作, 当定时器到期时会向一个通道发送一个值, 我们通过这个通道来处理这个事件
	timer := time.NewTimer(5 * time.Second) // 定义了一个两秒后触发的定时器

	<-timer.C // 等待定时器通道的信号

	fmt.Println("定时器触发")

	// 停止定时器, 返回boolen, true: 成功停止, false: 已经触发或者已经停止过了
	defer timer.Stop()

	// 重置定时器
	timer.Reset(3 * time.Second) // 重置为三秒触发

	// 2. go的time.Ticker: 可以定义定期触发任务, 而time.Timer只执行一次

	// 创建Ticket
	ticket := time.NewTicker(5 * time.Second)
	defer ticket.Stop()
	for t := range ticket.C {

		fmt.Println("Tick: ", t)

	}

}
