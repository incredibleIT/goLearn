package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup // 挂载计数器

func f1() {
	defer wait.Done() // 每完成一个协程Done掉一个
	fmt.Println("f1 begin")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 end")
}

func f2() {
	defer wait.Done()
	fmt.Println("f2 begin")
	time.Sleep(3 * time.Second)
	fmt.Println("f2 end")
}

//func timer(f func()) int64 {
//	startTime := time.Now().Unix()
//	f()
//	endTime := time.Now().Unix()
//	return endTime - startTime
//}

func main() {

	startTime := time.Now().Unix()
	wait.Add(2) // 挂载两个
	go f1()
	go f2()
	wait.Wait() // 必须等待所有协程走完
	endTime := time.Now().Unix()

	fmt.Println(endTime - startTime)
}
