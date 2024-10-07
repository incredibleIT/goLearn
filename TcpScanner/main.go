package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var port []int // 开放的端口追加到切片
var wg sync.WaitGroup
var semaphore = make(chan struct{}, 100)

func isOpen(p int) {

	defer wg.Done()

	semaphore <- struct{}{} // 获取信号量

	defer func() { <-semaphore }() // 释放信号量

	_, err := net.DialTimeout("tcp", "127.0.0.1:"+strconv.Itoa(p), time.Millisecond*2)

	if err == nil {
		fmt.Println(strconv.Itoa(p) + "此端口开放")
		port = append(port, p)
	} else {
		fmt.Println(strconv.Itoa(p) + "此端口不开放")
	}
}

// TCP 端口扫描工具
func main() {

	for p := 1; p <= 8888; p++ {
		wg.Add(1)
		go isOpen(p)
	}

	wg.Wait()
	fmt.Println(port)
}
