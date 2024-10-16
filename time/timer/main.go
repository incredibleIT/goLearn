package main

import (
	"fmt"
	"time"
)

func main() {
	// go的time.Timer 可用来设置一次性的定时操作, 当定时器到期时会向一个通道发送一个值, 我们通过这个通道来处理这个事件
	timer := time.NewTimer(5 * time.Second) // 定义了一个两秒后触发的定时器

	<-timer.C // 等待定时器通道的信号

	fmt.Println("定时器触发")

	// 停止定时器
	timer.Stop()

	// 重置定时器
	timer.Reset(3 * time.Second) // 重置为三秒触发

}
