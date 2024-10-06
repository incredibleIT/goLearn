package main

import (
	"fmt"
	"strconv"
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

var Chan = make(chan int)

//var wg sync.WaitGroup

var mux sync.Mutex

func main() {
	//wg.Add(1000)
	//for i := 0; i < 1000; i++ {
	//	go add()
	//}
	//
	//wg.Wait()
	//fmt.Println(x)
	k := 1
	for i := 0; i < 1; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			mux.Lock()
			k++
			mux.Unlock()
			fmt.Println("hello" + strconv.Itoa(i))
		}()
	}

	wg.Wait()

}
