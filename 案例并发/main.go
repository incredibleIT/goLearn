package main

import (
	"fmt"
	"sync"
	"time"
)

func f(i int, wg *sync.WaitGroup) {
	fmt.Printf("%d start\n", i)

	time.Sleep(time.Second * 3)

	fmt.Printf("%d end\n", i)

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go f(i, &wg)
	}

	wg.Wait()

}
