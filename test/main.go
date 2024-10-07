package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	var data []int

	for port := 1; port <= 300; port++ {
		wg.Add(1)
		go func(port int) {

			defer wg.Done()
			_, err := net.DialTimeout("tcp", "127.0.0.1:"+strconv.Itoa(port), time.Millisecond*20)

			if err == nil {
				fmt.Println("请求成功")
				data = append(data, port)
			} else {
				fmt.Println("fail.....")
			}
		}(port)
	}

	wg.Wait()

	fmt.Println(data)

}
