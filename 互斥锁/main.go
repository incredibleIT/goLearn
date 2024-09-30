package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var lock sync.Mutex // 互斥锁

var x = 0

func add() {
	defer wg.Done()

	lock.Lock() // 加锁
	x++
	lock.Unlock() // 解锁
}

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add()
	}

	wg.Wait()
	fmt.Println(x)

}
